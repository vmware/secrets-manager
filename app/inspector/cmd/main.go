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
	"fmt"

	"github.com/spiffe/vsecm-sdk-go/sentry"

	"github.com/vmware/secrets-manager/core/constants/symbol"
	"github.com/vmware/secrets-manager/lib/system"
)

func main() {
	go system.KeepAlive()

	// Fetch the secret from the VSecM Safe.
	d, err := sentry.Fetch()
	if err != nil {
		fmt.Println("Err:", err.Error())
		return
	}

	if d.Data == "" {
		fmt.Println(symbol.Null)
		return
	}

	// d.Data is a serialized collection of VSecM secrets.
	fmt.Println(d.Data)
}
