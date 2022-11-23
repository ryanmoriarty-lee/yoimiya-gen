package config

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
	"path/filepath"
)

type YoimiyaConfigProvider struct{}

// Register registe a new function for make a service instance
func (provider *YoimiyaConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewYoimiyaConfig
}

// Boot will called when the service instantiate
func (provider *YoimiyaConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *YoimiyaConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *YoimiyaConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

/// Name define the name for this service
func (provider *YoimiyaConfigProvider) Name() string {
	return contract.ConfigKey
}
