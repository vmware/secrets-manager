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
	"fmt"

	// log "github.com/vmware/secrets-manager/v2/core/log/rpc"
	"github.com/spiffe/spike-sdk-go/log"

	"io"
	"net/http"
)

func respond(cid *string, r *http.Response) {
	const fName = "safe.respond"

	if r == nil {
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Log().Error(fName, "message", "Post: Problem closing request body.", "err", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Log().Error(fName, "message",
			"Post: Unable to read the response body from VSecM Safe.",
			"err", err.Error())
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
