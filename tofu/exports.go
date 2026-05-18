// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofu

// This file exports internal types and values needed by
// github.com/pulumi-proserv/pulumi-tool-terraform-migrate to build and walk
// the eval graph once and reuse scopes for multiple module instances.

// ContextPluginsType is the exported alias for the internal contextPlugins type.
type ContextPluginsType = contextPlugins

// ContextPlugins returns the plugin library from a Context, allowing
// external callers to construct an EvalGraphBuilder.
func (c *Context) ContextPlugins() *ContextPluginsType {
	return c.plugins
}

// WalkEval is the exported form of the walkEval constant.
const WalkEval = walkEval
