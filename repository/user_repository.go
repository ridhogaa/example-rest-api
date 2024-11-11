package repository

import (
	"context"
	"database/sql"
	"example-rest-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User, tx *sql.Tx) domain.User
}
