package repository

import (
	"context"
	"database/sql"
	"example-rest-api/helper"
	"example-rest-api/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, user domain.User, tx *sql.Tx) domain.User {
	SQL := "insert into users(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}
