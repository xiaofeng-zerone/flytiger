package service

import (
	"github.com/xiaofeng-zerone/flytiger/server/service/ddosmechanism"
)

type ServiceGroup struct {
	DdosMechanismServiceGroup ddosmechanism.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
