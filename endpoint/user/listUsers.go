
package user

import (
	"fmt"
	"context"
	. "lyang/model"

	"github.com/graniticio/granitic/v2/ws"
	"github.com/graniticio/granitic/logging"
)

type UserReader interface {
	ListUser() ([]User, error)
}

type ListUsersLogic struct {
	Log          logging.Logger
	DBManager    UserReader
}

func (ul *ListUsersLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	users, err := ul.DBManager.ListUser()

	if(err != nil) {
		fmt.Println("Error while reading data")
		fmt.Println(err.Error())
		res.HTTPStatus = 500
		return
	}
	res.Body = users
	res.HTTPStatus = 201
}
