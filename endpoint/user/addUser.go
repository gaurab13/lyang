
package user

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type AddUserLogic struct {
}

type AddUserRequest struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
}

func (gl *AddUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *AddUserRequest) {
	person := make(map[string]string)
	person["name"] = cb.Name
	person["age"] = cb.Age
	res.Body = person
	res.HTTPStatus = 201
}
