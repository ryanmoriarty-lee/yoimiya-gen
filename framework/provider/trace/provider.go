package trace

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
)

type YoimiyaTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *YoimiyaTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewYoimiyaTraceService
}

// Boot will called when the service instantiate
func (provider *YoimiyaTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *YoimiyaTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *YoimiyaTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *YoimiyaTraceProvider) Name() string {
	return contract.TraceKey
}
