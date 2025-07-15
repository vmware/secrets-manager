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
	"context"
	bootstrap2 "github.com/vmware/secrets-manager/v2/app/safe/internal/bootstrap"
	server "github.com/vmware/secrets-manager/v2/app/safe/internal/server/engine"
	"github.com/vmware/secrets-manager/v2/app/safe/internal/state/io"
	"github.com/vmware/secrets-manager/v2/app/safe/internal/state/secret/collection"
	"github.com/vmware/secrets-manager/v2/core/constants/env"
	"github.com/vmware/secrets-manager/v2/core/constants/key"
	"github.com/vmware/secrets-manager/v2/core/crypto"
	"github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	cEnv "github.com/vmware/secrets-manager/v2/core/env"
	"github.com/vmware/secrets-manager/v2/core/log/std"
	"github.com/vmware/secrets-manager/v2/core/probe"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	std.PrintEnvironmentInfo(&id, []string{
		string(env.AppVersion),
		string(env.VSecMLogLevel),
	})

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), key.CorrelationId, &id),
	)
	defer cancel()

	if cEnv.BackingStoreForSafe() == data.Postgres {
		go func() {
			std.InfoLn(&id, "Backing store is postgres.")
			std.InfoLn(&id, "Secrets will be stored in-memory "+
				"until the internal config is loaded.")

			safeConfig, err := bootstrap2.PollForConfig(id, ctx)
			if err != nil {
				std.FatalLn(&id,
					"Failed to retrieve VSecM Safe internal configuration", err.Error())
			}

			std.InfoLn(&id,
				"VSecM Safe internal configuration loaded. Initializing database.")

			err = io.InitDB(safeConfig.Config.DataSourceName)
			if err != nil {
				std.FatalLn(&id, "Failed to initialize database:", err)
				return
			}

			std.InfoLn(&id, "Database connection initialized.")

			// Persist secrets that have not been persisted yet to Postgres.

			errChan := make(chan error, 1)

			collection.Secrets.Range(func(key any, value any) bool {
				v := value.(data.SecretStored)

				io.PersistToPostgres(v, errChan)

				// This will not block since the channel has a buffer of 1.
				for err := range errChan {
					if err != nil {
						std.ErrorLn(&id, "Error persisting secret", err.Error())
					}
				}

				return true
			})

			// Drain any remaining errors from the channel
			close(errChan)
			for err := range errChan {
				if err != nil {
					std.ErrorLn(&id, "Error persisting secret", err.Error())
				}
			}
		}()
	}

	std.InfoLn(&id, "Acquiring identity...")

	// Channel to notify when the bootstrap timeout has been reached.
	timedOut := make(chan bool, 1)

	// These channels must complete in a timely manner, otherwise
	// the timeOut will be fired and will crash the app.
	acquiredSvid := make(chan bool, 1)
	updatedSecret := make(chan bool, 1)
	serverStarted := make(chan bool, 1)

	// Monitor the progress of acquiring an identity, updating the age key,
	// and starting the server. If the timeout occurs before all three events
	// happen, the function logs a fatal message and the process crashes.
	go bootstrap2.Monitor(&id,
		bootstrap2.ChannelsToMonitor{
			AcquiredSvid:  acquiredSvid,
			UpdatedSecret: updatedSecret,
			ServerStarted: serverStarted,
		}, timedOut,
	)

	// Time out if things take too long.
	go bootstrap2.NotifyTimeout(timedOut)

	// Create initial cryptographic seeds off-cycle.
	go bootstrap2.CreateRootKey(&id, updatedSecret)

	// App is alive; however, not yet ready to accept connections.
	<-probe.CreateLiveness()

	std.InfoLn(&id, "before acquiring source...")
	source := bootstrap2.AcquireSource(ctx, acquiredSvid)
	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}

		// Close the source after the server (1) is done serving, likely
		// when the app is shutting down due to an eviction or a panic.
		if err := s.Close(); err != nil {
			std.InfoLn(&id, "Problem closing SVID Bundle source: %v\n", err)
		}
	}(source)

	// (1)
	if err := server.Serve(source, serverStarted); err != nil {
		std.FatalLn(&id, "failed to serve", err.Error())
	}
}
