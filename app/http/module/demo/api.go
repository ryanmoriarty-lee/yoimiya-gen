package demo

import (
	demoService "github.com/ryanmoriarty-lee/yoimiya-gen/app/provider/demo"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)

	//如果需要访问以下api，需要在该项目下的config/development设定好相关的配置，如果需要orm连接数据库或cache连接redis，请自行进行安装
	//r.GET("/demo/orm", api.DemoOrm)
	//r.GET("/demo/cache/redis", api.DemoRedis)
	//r.GET("/demo/cache/cache", api.DemoCache)

	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	c.JSON(200, "this is demo for dev all")
}

// Demo2  for godoc
// @Summary 获取所有学生
// @Description 获取所有学生,不进行分页
// @Produce  json
// @Tags demo
// @Success 200 {array} UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}
