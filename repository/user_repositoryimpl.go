package repository

import (
	"example-rest-api/helper"
	"example-rest-api/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (userRepository *UserRepositoryImpl) Save(user domain.User) domain.User {
	if err := userRepository.db.Save(&user); err != nil {
		helper.PanicIfError(err.Error)
	}
	return user
}

func (userRepository *UserRepositoryImpl) FindAll() []domain.User {
	var listUser []domain.User
	result := userRepository.db.Find(&listUser)
	if result != nil {
		helper.PanicIfError(result.Error)
	}
	return listUser
}
