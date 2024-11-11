package service

import (
	"example-rest-api/model/domain"
	"example-rest-api/model/request"
	"example-rest-api/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (userServiceImpl *UserServiceImpl) Save(request request.UserRequest) domain.User {

	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	user = userServiceImpl.UserRepository.Save(user)

	return user
}
