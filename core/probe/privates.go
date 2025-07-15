/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package probe

import (
	"fmt"
	"github.com/vmware/secrets-manager/v2/core/constants/val"
	"log"
	"net/http"
)

func ok(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, val.Ok)
	if err != nil {
		log.Printf("probe response failure: %s", err.Error())
		return
	}
}
