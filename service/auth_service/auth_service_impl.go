package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"go_share/helper"
	"go_share/model/api/api_request"
	"go_share/model/api/api_response"
	"go_share/model/domain"
	"go_share/repository/user_repository"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewAuthService(userRepository user_repository.UserRepository, DB *sql.DB, validator *validator.Validate) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validator: validator}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request api_request.AuthRegisterRequest) (api_response.AuthRegisterResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	channelRegister := make(chan domain.User, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		userExist, _ := service.UserRepository.FindByEmail(ctx, tx, request.Email)
		channelRegister <- userExist
	}()

	result := <-channelRegister

	// cek jika email ada
	empty := domain.User{}
	if result != empty {
		return api_response.AuthRegisterResponse{}, errors.New("Email telah terdaftar")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		helper.PanicIfError(err)
		tiketStr := strconv.Itoa(int(time.Now().Unix()))
		tiket := sql.NullString{String: tiketStr, Valid: true}
		result.Nama = request.Nama
		result.Email = request.Email
		result.IdLembaga = request.IdLembaga
		result.Tiket = tiket
		result.Password = string(hashedPassword)
		channelRegister <- result
	}()

	result = <-channelRegister

	wg.Add(1)
	go func() {
		defer wg.Done()
		user := service.UserRepository.Save(ctx, tx, result)
		channelRegister <- user
	}()

	result = <-channelRegister

	wg.Add(1)
	go func() {
		defer wg.Done()

		newUser, err := service.UserRepository.FindById(ctx, tx, result.IdUser)
		helper.PanicIfError(err)

		channelRegister <- newUser
	}()

	go func() {
		wg.Wait()
		close(channelRegister)
	}()

	result = <-channelRegister

	return api_response.AuthRegisterResponse{
		Email: result.Email,
		Tiket: result.Tiket,
	}, nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request api_request.AuthLoginRequest) (api_response.AuthLoginResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	channelUser := make(chan domain.User, 1)
	var wg sync.WaitGroup

	// cek email exist
	wg.Add(1)
	go func() {
		defer wg.Done()

		user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
		helper.PanicIfError(err)
		channelUser <- user
	}()

	result := <-channelUser

	// cek password
	wg.Add(1)
	go func() {
		defer wg.Done()

		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
		helper.PanicIfError(err)
	}()

	// create token
	wg.Add(1)
	go func() {
		defer wg.Done()

		token, err := helper.CreateToken(string(result.IdUser))
		helper.PanicIfError(err)
		result.Token = sql.NullString{String: token, Valid: true}

		channelUser <- result
	}()

	result = <-channelUser

	// update user
	wg.Add(1)
	go func() {
		defer wg.Done()

		user := service.UserRepository.Update(ctx, tx, result)
		channelUser <- user
	}()

	result = <-channelUser

	// find user
	wg.Add(1)
	go func() {
		defer wg.Done()

		user, err := service.UserRepository.FindById(ctx, tx, result.IdUser)
		helper.PanicIfError(err)

		channelUser <- user
	}()

	go func() {
		wg.Wait()
		close(channelUser)
	}()

	responseUser := <-channelUser

	return api_response.AuthLoginResponse{
		Id_User: responseUser.IdUser,
		Nama:    responseUser.Nama,
		Email:   responseUser.Email,
		Token:   responseUser.Token,
	}, nil
}
