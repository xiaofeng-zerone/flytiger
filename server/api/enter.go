package api

import "github.com/xiaofeng-zerone/flytiger/server/api/ddosmechanism"

type ApiGroup struct {
	DdosMechanismApiGroup ddosmechanism.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
