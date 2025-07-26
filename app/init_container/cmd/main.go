/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"os"

	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/system"

	"github.com/spiffe/vsecm-sdk-go/sentry"

	"github.com/vmware/secrets-manager/v2/core/env"
)

func initialized() bool {
	r, _ := sentry.Fetch()
	v := r.Data
	return v != ""
}

func bye() {
	os.Exit(0)
}

func main() {
	const fName = "init_container.main"

	log.Log().Info(fName, "message", "Starting VSecM Init Container")

	// Wait for a specified duration before exiting the init container.
	// This can be useful when you want things to reconcile before
	// starting the main container.
	go system.Watch(system.WatchConfig{
		WaitTimeBeforeExit:      env.WaitBeforeExitForInitContainer(),
		PollInterval:            env.PollIntervalForInitContainer(),
		InitializationPredicate: initialized,
		ExitAction:              bye,
	})

	// Block the process from exiting but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	system.KeepAlive()
}
