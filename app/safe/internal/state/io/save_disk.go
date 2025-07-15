/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"encoding/json"
	"errors"
	crypto2 "github.com/vmware/secrets-manager/v2/core/crypto"
	entity "github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	log "github.com/vmware/secrets-manager/v2/core/log/std"
	"io"
	"os"
	"sync"
)

var lastBackedUpIndex = make(map[string]int)

// Only one thread reaches lastBackupIndex at a time;
// but still using this lock for defensive programming.
var lastBackupIndexLock = sync.Mutex{}

func saveSecretToDisk(secret entity.SecretStored, dataPath string) error {
	data, err := json.Marshal(secret)
	if err != nil {
		return errors.Join(
			err,
			errors.New("saveSecretToDisk: failed to marshal secret"),
		)
	}

	file, err := os.Create(dataPath)
	if err != nil {
		return errors.Join(
			err,
			errors.New("saveSecretToDisk: failed to create file"),
		)
	}
	defer func(f io.ReadCloser) {
		err := f.Close()
		if err != nil {
			id := crypto2.Id()
			log.InfoLn(&id, "saveSecretToDisk: problem closing file", err.Error())
		}
	}(file)

	return crypto2.EncryptToWriterAes(file, string(data))
}
