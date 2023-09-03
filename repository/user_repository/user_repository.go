package user_repository

import (
	"context"
	"database/sql"
	"go_share/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
