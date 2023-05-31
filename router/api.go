package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-base/internal/controller"
)

// Api 业务相关功能的路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	//绑定路由
	group.Bind(
		controller.User,
	)
	//自定义路由
	group.GET("/demo/:fielda", controller.Demo.Get)
	group.POST("/demo", controller.Demo.Create)
	group.PUT("/demo/:id", controller.Demo.Update)
	group.DELETE("/demo/:id", controller.Demo.Delete)
}
