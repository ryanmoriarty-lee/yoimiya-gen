package app

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
)

// YoimiyaAppProvider 提供App的具体实现方法
type YoimiyaAppProvider struct {
	BaseFolder string
}

// Register 注册YoimiyaApp方法
func (h *YoimiyaAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewYoimiyaApp
}

// Boot 启动调用
func (h *YoimiyaAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *YoimiyaAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *YoimiyaAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *YoimiyaAppProvider) Name() string {
	return contract.AppKey
}
