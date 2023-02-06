package services

import (
	"github.com/prosline/micro-services/mvc/domain"
	"github.com/prosline/micro-services/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
