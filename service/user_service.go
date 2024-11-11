package service

import (
	"example-rest-api/model/domain"
	"example-rest-api/model/request"
)

type UserService interface {
	Save(request request.UserRequest) domain.User
}
