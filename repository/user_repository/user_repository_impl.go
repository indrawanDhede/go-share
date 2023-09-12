package user_repository

import (
	"context"
	"go_share/helper"
	"go_share/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	result := tx.Create(&domain.User{
		Nama:      user.Nama,
		Email:     user.Email,
		Password:  user.Password,
		IdLembaga: user.IdLembaga,
		Tiket:     user.Tiket,
		IsLogin:   user.IsLogin,
	}).Last(&user)

	helper.PanicIfError(result.Error)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.Model(domain.User{}).Where("id_user = ?", user.ID).Updates(domain.User{Nama: user.Nama, IdLembaga: user.IdLembaga, Token: user.Token, NoHp: user.NoHp, JenjangPendidikan: user.JenjangPendidikan, Bahasa: user.Bahasa, Alamat: user.Alamat, Kompetensi: user.Kompetensi}).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, id int) {
	user := domain.User{}
	err := tx.Where("id_users = ?", id).Delete(&user).Error
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, id int) (domain.User, error) {
	users := domain.User{}
	result := tx.Joins("Lembaga").Where("id_user = ?", id).First(&users)

	if result.Error != nil {
		return domain.User{}, result.Error
	} else {
		return users, nil
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {

	users := domain.User{}
	result := tx.Joins("Lembaga").Where("email = ?", email).First(&users)

	if result.Error != nil {
		return domain.User{}, result.Error
	} else {
		return users, nil
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	var users []domain.User
	tx.Find(&users)

	return users
}
