package response

import "github.com/xiaofeng-zerone/flytiger/server/model/ddosmechanism"

type HelloResponse struct {
	Hello ddosmechanism.Hello `json:"hello"`
}

type HelloListResponse struct {
	Hellos []ddosmechanism.Hello `json:"hellos"`
}
