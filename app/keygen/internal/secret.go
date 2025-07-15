/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package internal

import (
	"encoding/json"
	entity "github.com/vmware/secrets-manager/v2/core/entity/v1/reqres/safe"
	"github.com/vmware/secrets-manager/v2/core/env"
	"log"
	"os"
)

func secrets() entity.SecretEncryptedListResponse {
	p := env.ExportedSecretPathForKeyGen()

	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var secrets entity.SecretEncryptedListResponse

	err = json.Unmarshal(content, &secrets)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return secrets
}
