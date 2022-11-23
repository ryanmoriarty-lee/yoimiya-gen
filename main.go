package main

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/app/console"
	"github.com/ryanmoriarty-lee/yoimiya-gen/app/http"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/app"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/config"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/distributed"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/env"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/id"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/kernel"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/log"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewYoimiyaContainer()
	// 绑定App服务提供者
	container.Bind(&app.YoimiyaAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.YoimiyaEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.YoimiyaConfigProvider{})
	container.Bind(&id.YoimiyaIDProvider{})
	container.Bind(&trace.YoimiyaTraceProvider{})
	container.Bind(&log.YoimiyaLogServiceProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	engine, _ := http.NewHttpEngine(container)
	container.Bind(&kernel.YoimiyaKernelProvider{HttpEngine: engine})

	// 运行root命令
	console.RunCommand(container)
}
