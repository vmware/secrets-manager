/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"context"

	safe2 "github.com/vmware/secrets-manager/v2/app/sentinel/internal/safe"
	"github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	env2 "github.com/vmware/secrets-manager/v2/core/env"
	"github.com/vmware/secrets-manager/v2/core/spiffe"
	"os"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

// OSFileOpener is a struct that provides a method to open files.
type OSFileOpener struct{}

// Open opens the named file and returns an *os.File.
func (OSFileOpener) Open(name string) (*os.File, error) {
	return os.Open(name)
}

// EnvConfigReader is a struct that provides methods to read environment
// configurations.
type EnvConfigReader struct{}

// InitCommandPathForSentinel returns the path for the sentinel
// initialization command.
func (EnvConfigReader) InitCommandPathForSentinel() string {
	return env2.InitCommandPathForSentinel()
}

// InitCommandRunnerWaitBeforeExecIntervalForSentinel returns the wait interval
// before executing the sentinel initialization command.
func (EnvConfigReader) InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration {
	return env2.InitCommandRunnerWaitBeforeExecIntervalForSentinel()
}

// InitCommandRunnerWaitIntervalBeforeInitComplete returns the wait interval
// before the initialization is complete.
func (EnvConfigReader) InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration {
	return env2.InitCommandRunnerWaitIntervalBeforeInitComplete()
}

// NamespaceForVSecMSystem returns the namespace for the VSecM system.
func (EnvConfigReader) NamespaceForVSecMSystem() string {
	return env2.NamespaceForVSecMSystem()
}

// TODO: some of these entities do not belong here; move them elsewhere.

// SafeClient is a struct that provides methods to interact with the safe.
type SafeClient struct{}

// Check performs a check operation using the provided context and X509Source.
func (SafeClient) Check(ctx context.Context,
	src *workloadapi.X509Source) (int, string, error) {
	return safe2.Check(ctx, src)
}

// CheckInitialization checks the initialization status using the provided
// context and X509Source.
func (SafeClient) CheckInitialization(ctx context.Context,
	src *workloadapi.X509Source) (bool, error) {
	return safe2.CheckInitialization(ctx, src)
}

// Post sends a SentinelCommand to the safe using the provided context.
func (SafeClient) Post(ctx context.Context, sc data.SentinelCommand) error {
	return safe2.Post(ctx, sc)
}

// SpiffeClient is a struct that provides methods to interact with SPIFFE.
type SpiffeClient struct{}

// AcquireSourceForSentinel acquires an X509Source for the sentinel using
// the provided context.
func (SpiffeClient) AcquireSourceForSentinel(
	ctx context.Context) (*workloadapi.X509Source, bool) {
	return spiffe.AcquireSourceForSentinel(ctx)
}
