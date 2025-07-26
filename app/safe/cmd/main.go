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
	"github.com/spiffe/spike-sdk-go/log"

	"github.com/vmware/secrets-manager/v2/core/constants/env"
)

func main() {
	const fName = "safe.main"

	// TODO: any decryption/encryption needed can be done through SPIKE's APIs.

	// TODO: if we need a random seed, SPIKE SDK go has the functionality.

	// TODO: for now, v2 will only support Go SDK.

	log.Log().Info(fName, "app", env.AppVersion)

	log.Log().Info("msg", "Acquiring identity...")

	// TODO: acquire source from SPIKE SDK go.

	//// (1)
	//if err := server.Serve(source, serverStarted); err != nil {
	//	std.FatalLn(&id, "failed to serve", err.Error())
	//}
}
