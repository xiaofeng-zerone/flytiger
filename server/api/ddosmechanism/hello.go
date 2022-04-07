package ddosmechanism

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaofeng-zerone/flytiger/server/model/common/response"
	"github.com/xiaofeng-zerone/flytiger/server/model/ddosmechanism/request"
	helloRes "github.com/xiaofeng-zerone/flytiger/server/model/ddosmechanism/response"
)

type HelloApi struct{}

func (this *HelloApi) GetAll(c *gin.Context) {
	if err, hellos := helloService.GetAllHellos(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(helloRes.HelloListResponse{Hellos: hellos}, "获取成功", c)
	}
}

func (this *HelloApi) GetList(c *gin.Context) {
	var pageInfo request.SearchHelloParams
	_ = c.ShouldBindJSON(&pageInfo)
	//TODO-valid
	if err, list, total := helloService.GetListHellos(pageInfo.Hello, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
