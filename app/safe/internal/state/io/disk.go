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
	"github.com/vmware/secrets-manager/v2/core/constants/file"
	"github.com/vmware/secrets-manager/v2/core/constants/symbol"
	"github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	env2 "github.com/vmware/secrets-manager/v2/core/env"
	"math"
	"path"
	"strconv"

	"github.com/vmware/secrets-manager/v2/lib/backoff"
)

// PersistToDisk saves a given secret to disk and also creates a backup copy
// of the secret. The function is designed to enhance data durability through
// retries and backup management based on environmental configurations.
//
// Parameters:
//   - secret (entity.SecretStored): The secret to be saved, which is a
//     structured entity containing the secret's name and possibly other
//     metadata or the secret data itself.
//   - errChan (chan<- error): A channel through which errors are reported. This
//     channel allows the function to operate asynchronously, notifying the
//     caller of any issues in the process of persisting the secret.
func PersistToDisk(secret data.SecretStored, errChan chan<- error) {
	if env2.BackingStoreForSafe() != data.File {
		panic("Attempted to save to disk when backing store is not file")
	}

	backupCount := env2.SecretBackupCountForSafe()

	// Save the secret
	dataPath := path.Join(env2.DataPathForSafe(), secret.Name+file.AgeExtension)

	err := backoff.RetryExponential("PersistToDisk", func() error {
		return saveSecretToDisk(secret, dataPath)
	})

	if err != nil {
		errChan <- err
		// Do not proceed, since the primary save was not successful.
		return
	}

	lastBackupIndexLock.Lock()
	index, found := lastBackedUpIndex[secret.Name]
	if !found {
		lastBackedUpIndex[secret.Name] = 0
		index = 0
	}
	lastBackupIndexLock.Unlock()

	newIndex := math.Mod(float64(index+1), float64(backupCount))

	// Save a copy
	dataPath = path.Join(
		env2.DataPathForSafe(),
		secret.Name+symbol.FileNameSectionDelimiter+
			strconv.Itoa(int(newIndex))+
			symbol.FileNameSectionDelimiter+file.AgeBackupExtension,
	)

	err = backoff.RetryExponential(
		"PersistBackupToDisk", func() error {
			return saveSecretToDisk(secret, dataPath)
		})

	if err != nil {
		errChan <- err
		// Do not change lastBackedUpIndex
		// since the backup was not successful.
		return
	}

	lastBackupIndexLock.Lock()
	lastBackedUpIndex[secret.Name] = int(newIndex)
	lastBackupIndexLock.Unlock()
}
