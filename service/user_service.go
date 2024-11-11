package service

import (
	"context"
	"example-rest-api/model/domain"
	"example-rest-api/model/request"
)

type UserService interface {
	Save(ctx context.Context, request request.UserRequest) domain.User
}
