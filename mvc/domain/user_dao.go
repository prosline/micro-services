package domain

import (
	"fmt"
	"net/http"

	"github.com/prosline/micro-services/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Marcio", LastName: "DaSilva", Email: "marchito@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
