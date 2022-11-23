// Package http API.
// @title Yoimiya
// @version 0.0.1
// @description Yoimiya测试
// @termsOfService https://github.com/swaggo/swag

// @contact.name ryanmoriarty
// @contact.email 1255988688@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

import (
	_ "github.com/ryanmoriarty-lee/yoimiya-gen/app/http/swagger"
)
