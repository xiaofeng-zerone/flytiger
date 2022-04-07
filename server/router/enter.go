package router

import (
	"github.com/xiaofeng-zerone/flytiger/server/router/ddosmechanism"
)

type RouterGroup struct {
	DdosMechanismRouterGroup ddosmechanism.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
