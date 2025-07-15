/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package queue

import (
	"github.com/vmware/secrets-manager/v2/app/safe/internal/state/secret/queue/deletion"
	insertion2 "github.com/vmware/secrets-manager/v2/app/safe/internal/state/secret/queue/insertion"
)

// Initialize starts two goroutines: one to process the secret queue and
// another to process the Kubernetes secret queue. These goroutines are
// responsible for handling queued secrets and persisting them to disk.
func Initialize() {
	go insertion2.ProcessSecretBackingStoreQueue()
	go insertion2.ProcessK8sPrefixedSecretQueue()
	go deletion.ProcessSecretBackingStoreQueue()
}
