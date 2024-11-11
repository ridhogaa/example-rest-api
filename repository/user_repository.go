package repository

import (
	"example-rest-api/model/domain"
)

type UserRepository interface {
	Save(user domain.User) domain.User
}
