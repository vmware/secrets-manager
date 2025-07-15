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
	"fmt"
	cli2 "github.com/vmware/secrets-manager/v2/app/sentinel/internal/cli"
	safe2 "github.com/vmware/secrets-manager/v2/app/sentinel/internal/safe"
	"github.com/vmware/secrets-manager/v2/core/constants/env"
	"github.com/vmware/secrets-manager/v2/core/constants/key"
	"github.com/vmware/secrets-manager/v2/core/constants/sentinel"
	"github.com/vmware/secrets-manager/v2/core/crypto"
	entity "github.com/vmware/secrets-manager/v2/core/entity/v1/data"
	"os"
	"os/signal"
	"syscall"

	"github.com/akamensky/argparse"
)

func main() {
	id := crypto.Id()

	parser := argparse.NewParser(
		sentinel.CmdName,
		"Assigns secrets to workloads.",
	)

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(),
			key.CorrelationId, &id),
	)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case <-c:
			fmt.Println("Operation was cancelled.")
			// It is okay to cancel a cancelled context.
			cancel()
		}
	}()

	list := cli2.ParseList(parser)
	deleteSecret := cli2.ParseDeleteSecret(parser)
	namespaces := cli2.ParseNamespaces(parser)
	inputKeys := cli2.ParseInputKeys(parser)
	workloadIds := cli2.ParseWorkload(parser)
	secret := cli2.ParseSecret(parser)
	template := cli2.ParseTemplate(parser)
	format := cli2.ParseFormat(parser)
	encrypt := cli2.ParseEncrypt(parser)
	notBefore := cli2.ParseNotBefore(parser)
	expires := cli2.ParseExpires(parser)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		cli2.PrintUsage(parser)
		return
	}

	if *list {
		if *encrypt {
			err = safe2.Get(ctx, true)
			if err != nil {
				fmt.Println("Error getting from VSecM Safe:", err.Error())
				return
			}

			return
		}

		err = safe2.Get(ctx, false)
		if err != nil {
			fmt.Println("Error getting from VSecM Safe:", err.Error())
			return
		}

		return
	}

	if *namespaces == nil || len(*namespaces) == 0 {
		*namespaces = []string{string(env.Default)}
	}

	if cli2.InputValidationFailure(
		workloadIds, encrypt, inputKeys, secret, deleteSecret,
	) {
		return
	}

	err = safe2.Post(ctx, entity.SentinelCommand{
		WorkloadIds:        *workloadIds,
		Secret:             *secret,
		Namespaces:         *namespaces,
		Template:           *template,
		Format:             *format,
		Encrypt:            *encrypt,
		DeleteSecret:       *deleteSecret,
		SerializedRootKeys: *inputKeys,
		NotBefore:          *notBefore,
		Expires:            *expires,
	})

	if err != nil {
		fmt.Println("Error posting to VSecM Safe:", err.Error())
		return
	}
}
