// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalref

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/pulumi/opentofu/addrs"
	"github.com/pulumi/opentofu/configs"
	"github.com/pulumi/opentofu/configs/configload"
	"github.com/pulumi/opentofu/configs/configschema"
	"github.com/pulumi/opentofu/initwd"
	"github.com/pulumi/opentofu/providers"
	"github.com/pulumi/opentofu/registry"
)

func testAnalyzer(t *testing.T, fixtureName string) *Analyzer {
	configDir := filepath.Join("testdata", fixtureName)

	loader := configload.NewLoaderForTests(t)
	inst := initwd.NewModuleInstaller(loader.ModulesDir(), loader, registry.NewClient(t.Context(), nil, nil), nil)
	_, instDiags := inst.InstallModules(context.Background(), configDir, "tests", true, false, initwd.ModuleInstallHooksImpl{}, configs.RootModuleCallForTesting())
	if instDiags.HasErrors() {
		t.Fatalf("unexpected module installation errors: %s", instDiags.Err().Error())
	}
	if err := loader.RefreshModules(); err != nil {
		t.Fatalf("failed to refresh modules after install: %s", err)
	}

	cfg, loadDiags := loader.LoadConfig(t.Context(), configDir, configs.RootModuleCallForTesting())
	if loadDiags.HasErrors() {
		t.Fatalf("unexpected configuration errors: %s", loadDiags.Error())
	}

	resourceTypeSchema := &configschema.Block{
		Attributes: map[string]*configschema.Attribute{
			"string": {Type: cty.String, Optional: true},
			"number": {Type: cty.Number, Optional: true},
			"any":    {Type: cty.DynamicPseudoType, Optional: true},
		},
		BlockTypes: map[string]*configschema.NestedBlock{
			"single": {
				Nesting: configschema.NestingSingle,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"group": {
				Nesting: configschema.NestingGroup,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"list": {
				Nesting: configschema.NestingList,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"map": {
				Nesting: configschema.NestingMap,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"set": {
				Nesting: configschema.NestingSet,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
		},
	}
	schemas := map[addrs.Provider]providers.ProviderSchema{
		addrs.MustParseProviderSourceString("hashicorp/test"): {
			ResourceTypes: map[string]providers.Schema{
				"test_thing": {
					Block: resourceTypeSchema,
				},
			},
			DataSources: map[string]providers.Schema{
				"test_thing": {
					Block: resourceTypeSchema,
				},
			},
		},
	}

	return NewAnalyzer(cfg, schemas)
}
