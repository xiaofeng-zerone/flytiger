package request

import (
	"github.com/xiaofeng-zerone/flytiger/server/model/common/request"
	"github.com/xiaofeng-zerone/flytiger/server/model/ddosmechanism"
)

// hello分页条件查询及排序结构体
type SearchHelloParams struct {
	ddosmechanism.Hello
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
