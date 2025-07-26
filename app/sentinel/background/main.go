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
// "context"
// //server "github.com/vmware/secrets-manager/v2/app/safe/internal/server/engine"
// "github.com/vmware/secrets-manager/v2/app/sentinel/internal/initialization"
// e "github.com/vmware/secrets-manager/v2/core/constants/env"
// "github.com/vmware/secrets-manager/v2/core/constants/key"
// "github.com/vmware/secrets-manager/v2/core/crypto"
// //"github.com/vmware/secrets-manager/v2/core/env"
// "github.com/vmware/secrets-manager/v2/core/log/rpc"
// "github.com/vmware/secrets-manager/v2/core/log/std"
// "github.com/vmware/secrets-manager/v2/core/probe"
//
// // TODO: we don't need OIDC server functionality; we will need an OIDC client probably.
// // "github.com/vmware/secrets-manager/v2/app/sentinel/internal/oidc/server"
// "github.com/vmware/secrets-manager/v2/lib/system"
)

func main() {
	panic("implement me")

	//id := crypto.Id()
	//
	////Print the diagnostic information about the environment.
	//std.PrintEnvironmentInfo(&id, []string{
	//	string(e.AppVersion),
	//	string(e.VSecMLogLevel),
	//})
	//
	//<-probe.CreateLiveness()
	//go rpc.CreateLogServer()
	//
	//std.InfoLn(&id, "Executing the initialization commands (if any)")
	//
	//ctx := context.WithValue(context.Background(),
	//	key.CorrelationId, &id)
	//
	//std.TraceLn(&id, "before RunInitCommands")
	//
	//// Create the Initializer with all dependencies and run init commands
	//initialization.NewDefaultInitializer().RunInitCommands(ctx)
	//
	//std.InfoLn(&id, "Initialization commands executed successfully")
	//
	//// TODO: remove this env var.
	////if env.SentinelEnableOIDCResourceServer() {
	////	go server.Serve()
	////}
	//
	//// Run on the main thread to wait forever.
	//system.KeepAlive()
}
