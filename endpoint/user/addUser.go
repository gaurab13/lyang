
package user

import (
	"fmt"
	"context"

	"lyang/model"
	"github.com/graniticio/granitic/v2/ws"
	"github.com/graniticio/granitic/logging"
)

type UserCreator interface {
	AddUser(model.User) error
}

type AddUserLogic struct {
	Log          logging.Logger
	DBManager    UserCreator
}

type AddUserRequest struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
}

func (ul *AddUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *AddUserRequest) {
	person := make(map[string]interface{})
	person["name"] = cb.Name
	person["age"] = cb.Age
	person["gender"] = cb.Gender

	u := model.User{
		Name: cb.Name,
		Age: cb.Age,
		Gender: cb.Gender,
	}

	err := ul.DBManager.AddUser(u)

	if(err != nil) {
		fmt.Println("Error while writing to db")
		fmt.Println(err.Error())
		res.HTTPStatus = 500
		return
	}
	res.Body = person
	res.HTTPStatus = 201
}
