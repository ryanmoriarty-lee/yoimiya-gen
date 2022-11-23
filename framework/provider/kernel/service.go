package kernel

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/gin"
	"net/http"
)

// 引擎服务
type YoimiyaKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewYoimiyaKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &YoimiyaKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *YoimiyaKernelService) HttpEngine() http.Handler {
	return s.engine
}
