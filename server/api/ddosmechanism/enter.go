package ddosmechanism

import (
	"github.com/xiaofeng-zerone/flytiger/server/service"
)

type ApiGroup struct {
	HelloApi
}

var (
	helloService = service.ServiceGroupApp.DdosMechanismServiceGroup.HelloService
)
