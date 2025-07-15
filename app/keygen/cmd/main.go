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
	internal2 "github.com/vmware/secrets-manager/v2/app/keygen/internal"
	e "github.com/vmware/secrets-manager/v2/core/constants/env"
	crypto2 "github.com/vmware/secrets-manager/v2/core/crypto"
	"github.com/vmware/secrets-manager/v2/core/env"
	"github.com/vmware/secrets-manager/v2/core/log/std"
	"os"
)

func main() {
	id := crypto2.Id()

	// Print the diagnostic information about the environment.
	std.PrintEnvironmentInfo(&id, []string{
		string(e.AppVersion),
		string(e.VSecMLogLevel),
		string(e.VSecMKeygenDecrypt),
	})

	if env.KeyGenDecrypt() {
		// This is a Kubernetes Secret, mounted as a file.
		keyPath := env.RootKeyPathForKeyGen()

		if _, err := os.Stat(keyPath); os.IsNotExist(err) {
			std.FatalLn(&id,
				"CreateRootKey: Secret key not mounted at", keyPath)
			return
		}

		data, err := os.ReadFile(keyPath)
		if err != nil {
			std.FatalLn(&id,
				"CreateRootKey: Error reading file:", err.Error())
			return
		}

		// Root key needs to be committed to memory for VSecM Keygen to be able
		// to decrypt the secrets.
		secret := string(data)
		crypto2.SetRootKeyInMemory(secret)

		internal2.PrintDecryptedKeys()
		return
	}

	internal2.PrintGeneratedKeys()
}
