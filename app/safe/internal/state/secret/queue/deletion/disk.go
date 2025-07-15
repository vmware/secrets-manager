/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package deletion

import (
	"github.com/vmware/secrets-manager/v2/core/constants/file"
	"github.com/vmware/secrets-manager/v2/core/crypto"
	"github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	env2 "github.com/vmware/secrets-manager/v2/core/env"
	log "github.com/vmware/secrets-manager/v2/core/log/std"
	"os"
	"path"
)

// SecretDeleteQueue items are persisted to files. They are buffered, so that
// they can be written in the order they are queued and there are no concurrent
// writes to the same file at a time.
var SecretDeleteQueue = make(
	chan data.SecretStored,
	env2.SecretDeleteBufferSizeForSafe(),
)

// ProcessSecretBackingStoreQueue continuously processes a queue of secrets
// scheduled for deletion, removing each secret from disk. This function plays
// a crucial role in the secure management of secrets by ensuring that outdated
// or unnecessary secrets are not left stored, potentially posing a security risk.
//
// It operates in an endless loop, monitoring a global queue of secrets to be
// deleted.
func ProcessSecretBackingStoreQueue() {
	cid := crypto.Id()

	errChan := make(chan error)

	go func() {
		for e := range errChan {
			// If the `delete` operation spews out an error, log it.
			log.ErrorLn(&cid,
				"processSecretDeleteQueue: error deleting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(SecretDeleteQueue) == env2.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&cid,
				"processSecretDeleteQueue: "+
					"there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be removed from the disk.
		secret := <-SecretDeleteQueue

		store := env2.BackingStoreForSafe()
		switch store {
		case data.Memory:
			log.TraceLn(&cid, "ProcessSecretQueue: using in-memory store.")
			return
		case data.File:
			log.TraceLn(&cid, "ProcessSecretQueue: Will delete secret from disk.")
		case data.Kubernetes:
			panic("implement kubernetes store")
		case data.AwsSecretStore:
			panic("implement aws secret store")
		case data.AzureSecretStore:
			panic("implement azure secret store")
		case data.GcpSecretStore:
			panic("implement gcp secret store")
		case data.Postgres:
			log.WarnLn(&cid, "Delete operation has not been implemented for postgres backing store yet.")
			return
		}

		if secret.Name == "" {
			log.WarnLn(&cid,
				"processSecretDeleteQueue: trying to delete an empty secret. "+
					"Possibly picked a nil secret", len(SecretDeleteQueue))
			return
		}

		log.TraceLn(&cid,
			"processSecretDeleteQueue: picked a secret", len(SecretDeleteQueue))

		// Remove secret from disk.
		dataPath := path.Join(env2.DataPathForSafe(), secret.Name+file.AgeExtension)
		log.TraceLn(&cid,
			"processSecretDeleteQueue: removing secret from disk:", dataPath)
		err := os.Remove(dataPath)
		if err != nil && !os.IsNotExist(err) {
			log.WarnLn(&cid,
				"processSecretDeleteQueue: failed to remove secret", err.Error())
		}

		log.TraceLn(&cid,
			"processSecretDeleteQueue: should have deleted the secret.")
	}
}
