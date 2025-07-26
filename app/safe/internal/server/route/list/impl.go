/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package list

import (
	//"encoding/json"
	//"github.com/vmware/secrets-manager/v2/app/safe/internal/server/route/base/validation"
	//"github.com/vmware/secrets-manager/v2/app/safe/internal/state/secret/collection"
	//"github.com/vmware/secrets-manager/v2/core/audit/journal"
	//"github.com/vmware/secrets-manager/v2/core/constants/audit"
	//algo "github.com/vmware/secrets-manager/v2/core/constants/crypto"
	//"github.com/vmware/secrets-manager/v2/core/crypto"
	//"github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	//reqres "github.com/vmware/secrets-manager/v2/core/entity/v1/reqres/safe"
	//log "github.com/vmware/secrets-manager/v2/core/log/std"
	//"io"
	//"net/http"
	//
	//s "github.com/vmware/secrets-manager/v2/lib/spiffe"
	"net/http"
)

func doList(
	cid string, w http.ResponseWriter, r *http.Request, encrypted bool,
) {
	panic("implement me")

	//spiffeid := s.IdAsString(r)
	//
	//if !crypto.RootKeySetInMemory() {
	//	log.InfoLn(&cid, "Masked: Root key not set")
	//
	//	w.WriteHeader(http.StatusBadRequest)
	//	_, err := io.WriteString(w, "")
	//	if err != nil {
	//		log.InfoLn(&cid, "Masked: Problem with spiffeid", spiffeid)
	//	}
	//
	//	return
	//}
	//
	//// Only sentinel can list.
	//if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
	//	j.Event = audit.NotSentinel
	//	journal.Log(j)
	//	respond(w)
	//	return
	//}
	//
	//log.TraceLn(&cid, "Masked: before defer")
	//
	//defer func(b io.ReadCloser) {
	//	err := b.Close()
	//	if err != nil {
	//		log.InfoLn(&cid, "Masked: Problem closing body")
	//	}
	//}(r.Body)
	//
	//log.TraceLn(&cid, "Masked: after defer")
	//
	//secrets := collection.AllSecrets(cid)
	//
	//if encrypted {
	//	a := algo.Aes
	//
	//	secrets := collection.AllSecretsEncrypted(cid)
	//
	//	sfr := reqres.SecretEncryptedListResponse{
	//		Secrets:   secrets,
	//		Algorithm: a,
	//	}
	//
	//	j.Event = audit.Ok
	//	journal.Log(j)
	//
	//	resp, err := json.Marshal(sfr)
	//
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		_, err := io.WriteString(w, "Masked: Problem marshalling response")
	//		if err != nil {
	//			log.ErrorLn(&cid,
	//				"Masked: Problem sending response", err.Error())
	//		}
	//		return
	//	}
	//
	//	_, err = io.WriteString(w, string(resp))
	//	if err != nil {
	//		log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
	//	}
	//
	//	log.DebugLn(&cid, "Masked: after response")
	//	return
	//}
	//
	//sfr := reqres.SecretListResponse{
	//	Secrets: secrets,
	//}
	//
	//j.Event = audit.Ok
	//journal.Log(j)
	//
	//resp, err := json.Marshal(sfr)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	_, err := io.WriteString(w, "Masked: Problem marshalling response")
	//	if err != nil {
	//		log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
	//	}
	//	return
	//}
	//
	//_, err = io.WriteString(w, string(resp))
	//if err != nil {
	//	log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
	//}
	//
	//log.DebugLn(&cid, "Masked: after response")
}
