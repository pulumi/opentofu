package configs

import (
	"github.com/pulumi/opentofu/internal/configs"
	"github.com/spf13/afero"
)

type (
	Parser                = configs.Parser
	StaticModuleVariables = configs.StaticModuleVariables
	StaticModuleCall      = configs.StaticModuleCall
)

func NewParser(fs afero.Fs) *Parser {
	return configs.NewParser(fs)
}

func NewStaticModuleCall(
	addr []string, // addrs.Module
	vars StaticModuleVariables,
	rootPath string,
	workspace string,
) StaticModuleCall {
	return configs.NewStaticModuleCall(addr, vars, rootPath, workspace)
}
