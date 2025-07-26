/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package extract

import (
	"encoding/json"

	"github.com/spiffe/spike-sdk-go/log"

	entity "github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	"github.com/vmware/secrets-manager/v2/core/env"
	"regexp"
)

// WorkloadIdAndParts extracts the workload identifier and its constituent parts
// from a SPIFFE ID string, based on a predefined prefix that is removed from
// the SPIFFE ID.
//
// Parameters:
//   - spiffeid (string): The SPIFFE ID string from which the workload
//     identifier and parts are to be extracted.
//
// Returns:
//   - (string, []string): The first return value is the workload identifier,
//     which is essentially the first part of the SPIFFE ID after removing the
//     prefix. The second return value is a slice of strings representing all
//     parts of the SPIFFE ID after the prefix removal.
func WorkloadIdAndParts(spiffeid string) (string, []string) {
	re := env.NameRegExpForWorkload()
	if re == "" {
		return "", nil
	}
	wre := regexp.MustCompile(env.NameRegExpForWorkload())

	match := wre.FindStringSubmatch(spiffeid)

	if len(match) > 1 {
		return match[1], match
	}

	return "", nil
}

// SecretValue determines the appropriate representation of a secret's value to
// use, giving precedence to a transformed value over the raw values. This
// function supports both current implementations and backward compatibility by
// handling secrets with multiple values or transformed values.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - secret (*entity.SecretStored): A pointer to the secret entity from which
//     the value is to be retrieved.
//
// Returns:
//   - string: The selected representation of the secret's value. This could be
//     the transformed value, a single raw value, a JSON-encoded string of
//     multiple values, or an empty string in case of an error.
func SecretValue(cid string, secrets []entity.SecretStored) string {
	const fName = "Fetch:SecretValue"
	if secrets == nil {
		return ""
	}

	if len(secrets) == 0 {
		return ""
	}

	if len(secrets) == 1 {
		secret := secrets[0]

		if secret.ValueTransformed != "" {
			log.Log().Info(fName, "message", "Fetch: using transformed value")
			return secret.ValueTransformed
		}

		// This part is for backwards compatibility.
		// It probably won't execute because `secret.ValueTransformed` will
		// always be set.

		log.Log().Info(fName, "message", "Fetch: no transformed value found. returning raw value")
		return secret.Value
	}

	jsonData, err := json.Marshal(secrets)
	if err != nil {
		log.Log().Warn(fName, "message", "Fetch: Problem marshaling secrets", "err", err.Error())
		return ""
	}

	return string(jsonData)
}
