package xinit

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaofeng-zerone/flytiger/server/router"
)

//RouterInit 路由初始化
func RouterInit() *gin.Engine {
	Router := gin.Default()
	DdosMechanismRouter := router.RouterGroupApp.DdosMechanismRouterGroup

	PublicGroup := Router.Group("api")
	{
		DdosMechanismRouter.InitHelloRouter(PublicGroup) //注册Hello功能路由 不做鉴权
	}

	return Router
}
