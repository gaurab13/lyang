
package user

import (
	"context"
	"fmt"

	"lyang/model"
	"github.com/graniticio/granitic/v2/ws"
	"github.com/graniticio/granitic/logging"
)

type UserUpdater interface {
	UpdateUser(model.User) error
}

type UpdateUserLogic struct {
	Log          logging.Logger
	DBManager    UserUpdater
}

type UpdateUserRequest struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
}

func (ul *UpdateUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *UpdateUserRequest) {
	u := model.User{
		Name: cb.Name,
		Age: cb.Age,
		Gender: cb.Gender,
	}

	err := ul.DBManager.UpdateUser(u)

	if(err != nil) {
		fmt.Println("Error while updating to db")
		res.HTTPStatus = 500
		return
	}
	res.Body = u
	res.HTTPStatus = 201
}
