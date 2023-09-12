package user_service

import (
	"context"
	"github.com/go-playground/validator"
	"go_share/helper"
	"go_share/model/api/api_response"
	"go_share/model/domain"
	"go_share/repository/user_repository"
	"gorm.io/gorm"
	"sync"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *gorm.DB
	Validator      *validator.Validate
}

func NewUserService(userRepository user_repository.UserRepository, DB *gorm.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validator: validator}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []api_response.UserResponse {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	channelUser := make(chan []domain.User, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		result := service.UserRepository.FindAll(ctx, tx)
		channelUser <- result
	}()

	go func() {
		wg.Wait()
		close(channelUser)
	}()

	users := <-channelUser

	var responses []api_response.UserResponse
	for _, user := range users {
		response := api_response.UserResponse{
			Id_User:            user.IdUser,
			Nama:               user.Nama,
			Email:              user.Email,
			Token:              user.Token,
			Link_Foto:          user.LinkFoto,
			No_Hp:              user.NoHp,
			Jenjang_pendidikan: user.JenjangPendidikan,
			Alamat:             user.Alamat,
			Bahasa:             user.Bahasa,
			Kompetensi:         user.Kompetensi,
			Is_Login:           user.IsLogin,
		}
		responses = append(responses, response)
	}

	return responses
}
