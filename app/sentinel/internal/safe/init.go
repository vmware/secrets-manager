/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"context"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware/secrets-manager/v2/core/constants/key"
	"github.com/vmware/secrets-manager/v2/core/validation"
)

// CheckInitialization verifies if VSecM Sentinel has executed its init commands
// stanza successfully. This function utilizes a SPIFFE-based mTLS
// authentication mechanism to securely connect to a specified API endpoint.
//
// Parameters:
//   - ctx context.Context: The context carrying the correlation ID used for
//     logging and tracing the operation across different system components.
//     The correlation ID is extracted from the context for error logging
//     purposes.
//   - source *workloadapi.X509Source: A pointer to an X509Source, which
//     provides the credentials necessary for mTLS configuration. The source
//     must not be nil, as it is essential for establishing the TLS connection.
//
// Returns:
//   - bool: Returns true if VSecM Sentinel is initialized; false otherwise .
//   - error: Returns an error if the workload source is nil, URL joining fails,
//     the API call fails, the response body cannot be read, or the JSON
//     response cannot be unmarshalled. The error will provide a detailed
//     message about the nature of the failure.
func CheckInitialization(
	ctx context.Context, source *workloadapi.X509Source,
) (bool, error) {
	// TODO: get rid of all correlation ids.
	_ = ctx.Value(key.CorrelationId).(*string)

	if source == nil {
		return false, errors.New("check: workload source is nil")
	}

	// Authorizer
	_ = tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New(
			"I don't know you, and it's crazy: '" + id.String() + "'",
		)
	})

	panic("TODO: this needs an updated implementation. set a secret after init command execution; and verify that it is set here.")

	return false, nil
}
