// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/pulumi/opentofu/internal/grpcwrap"
	plugin "github.com/pulumi/opentofu/internal/plugin6"
	simple "github.com/pulumi/opentofu/internal/provider-simple-v6"
	"github.com/pulumi/opentofu/internal/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
