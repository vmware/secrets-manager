/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package secret

import (
	"io"
	"net/http"
	"strings"
	"time"

	net "github.com/vmware/secrets-manager/app/safe/internal/server/route/base/http"
	"github.com/vmware/secrets-manager/app/safe/internal/server/route/base/json"
	"github.com/vmware/secrets-manager/app/safe/internal/server/route/base/state"
	"github.com/vmware/secrets-manager/app/safe/internal/server/route/base/validation"
	ioState "github.com/vmware/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware/secrets-manager/core/audit/journal"
	"github.com/vmware/secrets-manager/core/constants/audit"
	"github.com/vmware/secrets-manager/core/constants/val"
	"github.com/vmware/secrets-manager/core/crypto"
	entity "github.com/vmware/secrets-manager/core/entity/v1/data"
	"github.com/vmware/secrets-manager/core/env"
	log "github.com/vmware/secrets-manager/core/log/std"
	data "github.com/vmware/secrets-manager/lib/entity"
	s "github.com/vmware/secrets-manager/lib/spiffe"
)

// Secret handles the creation, updating, and management of secrets.
// It performs several checks and operations based on the request parameters.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     logging and tracking purposes.
//   - w: An http.ResponseWriter object used to send responses back to the
//     client.
//   - r: An http.Request object containing the details of the client's request.
//   - spiffeid: A string representing the SPIFFE ID of the client making the
//     request.
func Secret(cid string, r *http.Request, w http.ResponseWriter) {
	spiffeid := s.IdAsString(r)
	if spiffeid == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, val.NotOk)
		if err != nil {
			log.ErrorLn(&cid, "error writing response", err.Error())
		}

		return
	}

	if !crypto.RootKeySetInMemory() {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, val.NotOk)
		if err != nil {
			log.ErrorLn(&cid, "error writing response", err.Error())
		}
		log.InfoLn(&cid, "Secret: Root key not set")

		return
	}

	j := journal.CreateDefaultEntry(cid, spiffeid, r)
	journal.Log(j)

	isSentinelOrScout, respond := validation.IsSentinelOrScout(j, cid, spiffeid)

	// Only sentinel or scout can do this
	if !isSentinelOrScout {
		j.Event = audit.NotSentinelOrScout
		journal.Log(j)
		respond(w)
		return
	}

	log.DebugLn(&cid, "Secret: sentinel spiffeid:", spiffeid)

	body, _ := net.ReadBody(cid, r)
	if body == nil {
		j.Event = audit.BadPayload
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	log.DebugLn(&cid, "Secret: Parsed request body")

	ur, _ := json.UnmarshalSecretUpsertRequest(body)
	if ur == nil {
		j.Event = audit.BadPayload
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	sr := *ur

	workloadIds := sr.WorkloadIds
	value := sr.Value
	namespaces := sr.Namespaces
	template := sr.Template
	format := sr.Format
	encrypt := sr.Encrypt
	notBefore := sr.NotBefore
	expiresAfter := sr.Expires

	// The next check is only for non-vsecm-safe workloads.
	if len(workloadIds) == 0 || workloadIds[0] != "vsecm-safe" {
		log.TraceLn(&cid, "Secret: workloadIds is empty or not vsecm-safe")

		// If postgres mode enabled and db is not initialized, return error.
		if env.BackingStoreForSafe() == entity.Postgres && !ioState.PostgresReady() {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, err := io.WriteString(w, val.NotOk)
			if err != nil {
				log.ErrorLn(&cid, "error writing response", err.Error())
			}
			log.InfoLn(&cid, "Secret: Database not initialized")
			return
		}
	} else {
		log.TraceLn(&cid, "Secret: vsecm-safe workload detected")
	}

	if len(workloadIds) == 0 && encrypt {
		isSentinel, respond := validation.IsSentinel(j, cid, spiffeid)
		if !isSentinel {
			j.Event = audit.NotSentinel
			journal.Log(j)
			respond(w)
		}

		net.SendEncryptedValue(cid, value, j, w)

		return
	}

	if len(namespaces) == 0 {
		namespaces = []string{"default"}
	}

	log.DebugLn(&cid, "Secret:Upsert: ", "workloadIds:", workloadIds,
		"namespaces:", namespaces,
		"template:", template, "format:", format, "encrypt:", encrypt,
		"notBefore:", notBefore, "expiresAfter:", expiresAfter)

	if len(workloadIds) == 0 && !encrypt {
		log.TraceLn(&cid, "Secret:Upsert: No workload id. Exiting")

		j.Event = audit.NoWorkloadId
		journal.Log(j)

		return
	}

	// `encrypt` means that the value is encrypted, so we need to decrypt it.
	if encrypt {
		log.TraceLn(&cid, "Secret: Value is encrypted")

		decrypted, err := crypto.DecryptValue(value)

		// If decryption failed, return an error response.
		if err != nil {
			log.InfoLn(&cid, "Secret: Decryption failed", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid,
					"Secret: Problem sending response", err.Error())
			}

			return
		}

		// Update the value of the request to the decoded value.
		sr.Value = decrypted
		value = sr.Value
	} else {
		log.TraceLn(&cid, "Secret: Value is not encrypted")
	}

	nb := data.JsonTime{}
	exp := data.JsonTime{}

	if notBefore == "now" {
		nb = data.JsonTime(time.Now())
	} else {
		nbTime, err := time.Parse(time.RFC3339, notBefore)
		if err != nil {
			nb = data.JsonTime(time.Now())
		} else {
			nb = data.JsonTime(nbTime)
		}
	}

	if expiresAfter == "never" {
		// This is the largest time go stdlib can represent.
		// It is far enough into the future that the author does not care
		// what happens after.
		exp = data.JsonTime(
			time.Date(9999, time.December,
				31, 23, 59, 59, 999999999, time.UTC),
		)
	} else {
		expTime, err := time.Parse(time.RFC3339, expiresAfter)
		if err != nil {
			exp = data.JsonTime(
				time.Date(9999, time.December,
					31, 23, 59, 59, 999999999, time.UTC),
			)
		} else {
			exp = data.JsonTime(expTime)
		}
	}

	for _, workloadId := range workloadIds {
		isClerk, _ := validation.IsClerk(j, cid, spiffeid)

		// Clerk can only set `raw:` secrets.
		if isClerk && !strings.HasPrefix(workloadId, "raw:") {
			log.WarnLn(&cid, "Clerk is trying to upsert non-raw secrets."+
				" Skipping:", workloadId)
			continue
		}

		secretToStore := entity.SecretStored{
			Name: workloadId,
			Meta: entity.SecretMeta{
				Namespaces:    namespaces,
				Template:      template,
				Format:        format,
				CorrelationId: cid,
			},
			Value:        value,
			NotBefore:    time.Time(nb),
			ExpiresAfter: time.Time(exp),
		}

		state.Upsert(secretToStore, workloadId, cid, j, w)
	}
}
