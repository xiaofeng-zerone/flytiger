package ddosmechanism

import (
	"github.com/xiaofeng-zerone/flytiger/server/model/common/request"
	"github.com/xiaofeng-zerone/flytiger/server/model/ddosmechanism"
)

type HelloService struct{}

var HelloServiceApp = new(HelloService)

func (helloService *HelloService) GetAllHellos() (err error, hellos []ddosmechanism.Hello) {
	//TODO-db
	allData := []ddosmechanism.Hello{
		{
			Name:   "xiao1",
			Remark: "test1 for all",
		},
		{
			Name:   "xiao2",
			Remark: "test2 for all",
		},
	}
	return nil, allData
}

func (helloService *HelloService) GetListHellos(hello ddosmechanism.Hello, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	//TODO-db
	listData := []ddosmechanism.Hello{
		{
			Name:   "xiao1",
			Remark: "test1 for list",
		},
		{
			Name:   "xiao2",
			Remark: "test2 for list",
		},
	}
	total = int64(len(listData))

	return nil, listData, total
}
