package service

import (
	"context"
	"database/sql"
	"example-rest-api/helper"
	"example-rest-api/model/domain"
	"example-rest-api/model/request"
	"example-rest-api/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (userServiceImpl *UserServiceImpl) Save(ctx context.Context, request request.UserRequest) domain.User {
	err := userServiceImpl.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	user = userServiceImpl.UserRepository.Save(ctx, user, tx)

	return user
}
