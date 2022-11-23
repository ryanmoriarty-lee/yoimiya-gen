package env

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
)

type YoimiyaEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *YoimiyaEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewYoimiyaEnv
}

// Boot will called when the service instantiate
func (provider *YoimiyaEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *YoimiyaEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *YoimiyaEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *YoimiyaEnvProvider) Name() string {
	return contract.EnvKey
}
