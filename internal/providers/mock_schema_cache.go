package providers

import "github.com/pulumi/opentofu/internal/addrs"

func NewMockSchemaCache() *schemaCache {
	return &schemaCache{
		m: make(map[addrs.Provider]ProviderSchema),
	}
}
