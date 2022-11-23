package id

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
)

type YoimiyaIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *YoimiyaIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewYoimiyaIDService
}

// Boot will called when the service instantiate
func (provider *YoimiyaIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *YoimiyaIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *YoimiyaIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *YoimiyaIDProvider) Name() string {
	return contract.IDKey
}
