package user_repository

import (
	"context"
	"go_share/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Delete(ctx context.Context, tx *gorm.DB, id int)
	FindById(ctx context.Context, tx *gorm.DB, id int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.User
}
