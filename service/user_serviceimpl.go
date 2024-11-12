package service

import (
	"example-rest-api/model/domain"
	"example-rest-api/model/request"
	"example-rest-api/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (u *UserServiceImpl) Save(request request.UserRequest) domain.User {

	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	user = u.UserRepository.Save(user)

	return user
}

func (u *UserServiceImpl) FindAll() []domain.User {
	return u.UserRepository.FindAll()
}
