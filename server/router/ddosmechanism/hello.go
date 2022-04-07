package ddosmechanism

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaofeng-zerone/flytiger/server/api"
)

type HelloRouter struct{}

func (s *HelloRouter) InitHelloRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	helloRouter := Router.Group("hello")
	helloApi := api.ApiGroupApp.DdosMechanismApiGroup.HelloApi

	{
		//curl  http://192.168.11.102:6666/api/hello/all
		helloRouter.GET("all", helloApi.GetAll) // 获取所有

		//curl -X POST -d '{"page": 1, "pageSize": 123}' -H 'Content-Type: application/json'  http://192.168.11.102:6666/api/hello/list
		helloRouter.POST("list", helloApi.GetList) // 获取列表

	}
	return helloRouter
}
