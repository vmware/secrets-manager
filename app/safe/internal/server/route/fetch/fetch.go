/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vmware/secrets-manager/app/safe/internal/server/route/base/extract"
	"github.com/vmware/secrets-manager/app/safe/internal/server/route/base/handle"
	rv "github.com/vmware/secrets-manager/app/safe/internal/server/route/base/validation"
	"github.com/vmware/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware/secrets-manager/core/audit/journal"
	"github.com/vmware/secrets-manager/core/constants/audit"
	"github.com/vmware/secrets-manager/core/crypto"
	"github.com/vmware/secrets-manager/core/entity/v1/data"
	reqres "github.com/vmware/secrets-manager/core/entity/v1/reqres/safe"
	log "github.com/vmware/secrets-manager/core/log/std"
	"github.com/vmware/secrets-manager/core/validation"
	s "github.com/vmware/secrets-manager/lib/spiffe"
)

// Fetch handles the retrieval of a secret for a given workload, identified by
// its SPIFFE ID.
// The function performs several checks to ensure the request is valid and then
// fetches the secret.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     tracking and logging purposes.
//   - w: An http.ResponseWriter object used to send responses back to the
//     client.
//   - r: An http.Request object containing the request details from the client.
//   - spiffeid: A string representing the SPIFFE ID of the client making the
//     request.
func Fetch(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	spiffeid := s.IdAsString(r)

	if !crypto.RootKeySetInMemory() {
		log.InfoLn(&cid, "Fetch: Root key not set")

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(
				&cid,
				"Status: problem sending response", spiffeid)
		}

		return
	}

	if !rv.CheckDatabaseReadiness(cid, w) {
		return
	}

	j := data.JournalEntry{
		CorrelationId: cid,
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.Enter,
	}

	journal.Log(j)

	// Only workloads can fetch.
	if !validation.IsWorkload(spiffeid) {
		handle.BadSvidResponse(cid, w, spiffeid, j)
		return
	}

	log.DebugLn(&cid, "Fetch: sending response")

	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem closing body")
		}
	}(r.Body)

	log.DebugLn(&cid, "Fetch: preparing request")

	workloadId, parts := extract.WorkloadIdAndParts(spiffeid)
	if len(parts) == 0 {
		handle.BadPeerSvidResponse(cid, w, spiffeid, j)
		return
	}

	var secrets []data.SecretStored

	if workloadId == "vsecm-scout" {
		secrets = collection.RawSecrets(cid)
	} else {
		secret, err := collection.ReadSecret(cid, workloadId)
		if err != nil {
			log.WarnLn(&cid, "Fetch: Attempted to read secret from disk.")
			log.TraceLn(&cid,
				"Likely expected error. No need to panic:", err.Error())
		}

		if secret != nil {
			secrets = append(secrets, *secret)
		}
	}

	log.TraceLn(&cid, "Fetch: workloadId", workloadId)

	// If secret does not exist, send an empty response.
	if secrets == nil {
		handle.NoSecretResponse(cid, w, j)
		return
	}

	if len(secrets) == 0 {
		handle.NoSecretResponse(cid, w, j)
		return
	}

	if workloadId == "vsecm-scout" {
		value := extract.SecretValue(cid, secrets)

		sfr := reqres.SecretFetchResponse{
			Data: value,
		}

		handle.SuccessResponse(cid, w, j, sfr)
		return
	}

	// Only vsecm-scout workloads can fetch multiple `raw` secrets.
	if len(secrets) > 1 {
		log.WarnLn(&cid, "Fetch: Sending 'no secrets' response:",
			workloadId, len(secrets))
		handle.NoSecretResponse(cid, w, j)
		return
	}

	// Regular workloads will only have one secret.
	secret := secrets[0]

	log.DebugLn(&cid, "Fetch: will send. workload id:", workloadId)

	value := extract.SecretValue(cid, secrets)

	// RFC3339 is what Go uses internally when marshaling dates.
	// Choosing it to be consistent.
	sfr := reqres.SecretFetchResponse{
		Data: value,
		Created: fmt.Sprintf("\"%s\"",
			secret.Created.Format(time.RFC3339)),
		Updated: fmt.Sprintf("\"%s\"",
			secret.Updated.Format(time.RFC3339)),
	}

	handle.SuccessResponse(cid, w, j, sfr)
}
