/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	e "github.com/vmware/secrets-manager/core/constants/env"
	"github.com/vmware/secrets-manager/core/env"
)

// Any SPIFFE ID regular expression matcher shall start with the
// `^spiffe://$trustDomain` prefix for extra security.
//
// This variable shall be treated as constant and should not be modified.
var spiffeRegexPrefixStart = "^spiffe://" + env.SpiffeTrustDomain() + "/"
var spiffeIdPrefixStart = "spiffe://" + env.SpiffeTrustDomain() + "/"

// IsWorkload checks if a given SPIFFE ID belongs to a workload.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to a workload, `false` otherwise.
func IsWorkload(spiffeid string) bool {
	prefix := env.SpiffeIdPrefixForWorkload()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixWorkload) +
					" environment variable. " +
					" val: " + env.SpiffeIdPrefixForWorkload() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		nrw := env.NameRegExpForWorkload()
		wre, err := regexp.Compile(nrw)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for SPIFFE ID." +
					" Check the " + string(e.VSecMWorkloadNameRegExp) +
					" environment variable." +
					" val: " + env.NameRegExpForWorkload() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		match := wre.FindStringSubmatch(spiffeid)
		if len(match) == 0 {
			return false
		}

		return re.MatchString(spiffeid)
	}

	if !strings.HasPrefix(spiffeid, spiffeIdPrefixStart) {
		return false
	}

	nrw := env.NameRegExpForWorkload()
	if !strings.HasPrefix(nrw, spiffeRegexPrefixStart) {

		// Insecure configuration detected.
		// Panic to prevent further issues:
		panic(
			"Invalid regular expression pattern for SPIFFE ID." +
				" Expected: ^spiffe://<trust_domain>/..." +
				" Check the " + string(e.VSecMWorkloadNameRegExp) +
				" environment variable." +
				" val: " + env.NameRegExpForWorkload() +
				" trust: " + env.SpiffeTrustDomain(),
		)
	}

	wre, err := regexp.Compile(nrw)
	if err != nil {
		panic(
			"Failed to compile the regular expression pattern " +
				"for SPIFFE ID." +
				" Check the " + string(e.VSecMWorkloadNameRegExp) +
				" environment variable." +
				" val: " + env.NameRegExpForWorkload() +
				" trust: " + env.SpiffeTrustDomain(),
		)
	}

	match := wre.FindStringSubmatch(spiffeid)
	if len(match) == 0 {
		return false
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// IsSentinel checks if a given SPIFFE ID belongs to VSecM Sentinel.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper  prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to VSecM Sentinel, `false` otherwise.
func IsSentinel(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForSentinel()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for VSecM Sentinel SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixSentinel) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForSentinel() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

func IsScout(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForScout()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for VSecM Scout SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixScout) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForScout() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

func IsClerk(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForClerk()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for VSecM Clerk SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixClerk) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForClerk() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// IsSafe checks if a given SPIFFE ID belongs to VSecM Safe.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to VSecM Safe, `false` otherwise.
func IsSafe(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForSafe()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for Sentinel SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixSafe) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForSafe() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

func IsRelayServer(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForRelayServer()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for Relay Server SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixRelayServer) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForRelayServer() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

func IsRelayClient(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForRelayClient()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for Relay Client SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixRelayClient) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForRelayClient() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// EnsureSafe checks the safety of the SPIFFE ID from the provided X509Source.
// It retrieves an X.509 SVID (SPIFFE Verifiable Identity Document) from the
// source, and validates the SPIFFE ID against a predefined safety check.
//
// If the X509Source fails to provide an SVID, the function will panic with an
// error message specifying the inability to retrieve the SVID.
//
// Similarly, if the SPIFFE ID from the retrieved SVID does not pass the safety
// check, the function will panic with an error message indicating that the
// SPIFFE ID is not recognized.
//
// Panicking in this function indicates severe issues with identity verification
// that require immediate attention and resolution.
//
// Usage:
//
//	var source *workloadapi.X509Source // Assume source is properly initialized
//	EnsureSafe(source)
func EnsureSafe(source *workloadapi.X509Source) {
	svid, err := source.GetX509SVID()
	if err != nil {
		panic(
			fmt.Sprintf(
				"Unable to get X.509 SVID from source bundle: %s",
				err.Error(),
			),
		)
	}

	svidId := svid.ID
	if !IsSafe(svidId.String()) {
		panic(
			fmt.Sprintf(
				"SpiffeId check: I don't know you, and it's crazy: %s",
				svidId.String(),
			),
		)
	}
}

func EnsureRelayServer(source *workloadapi.X509Source) {
	svid, err := source.GetX509SVID()
	if err != nil {
		panic(
			fmt.Sprintf(
				"Unable to get X.509 SVID from source bundle: %s",
				err.Error(),
			),
		)
	}

	svidId := svid.ID
	if !IsRelayServer(svidId.String()) {
		panic(
			fmt.Sprintf(
				"SpiffeId check: I don't know you, and it's crazy: %s",
				svidId.String(),
			),
		)
	}
}

func EnsureRelayClient(source *workloadapi.X509Source) {
	svid, err := source.GetX509SVID()
	if err != nil {
		panic(
			fmt.Sprintf(
				"Unable to get X.509 SVID from source bundle: %s",
				err.Error(),
			),
		)
	}

	svidId := svid.ID
	if !IsRelayClient(svidId.String()) {
		panic(
			fmt.Sprintf(
				"SpiffeId check: I don't know you, and it's crazy: %s",
				svidId.String(),
			),
		)
	}
}
