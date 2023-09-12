package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"go_share/helper"
	"go_share/model/api/api_request"
	"go_share/model/api/api_response"
	"go_share/model/domain"
	"go_share/repository/user_repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

type AuthServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *gorm.DB
	Validator      *validator.Validate
}

func NewAuthService(userRepository user_repository.UserRepository, DB *gorm.DB, validator *validator.Validate) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validator: validator}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request api_request.AuthRegisterRequest) (api_response.AuthRegisterResponse, error) {
	tx := service.DB.WithContext(ctx).Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	channelRegister := make(chan domain.User, 1)
	var wg sync.WaitGroup

	empty := domain.User{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		userExist, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
		if err == gorm.ErrRecordNotFound {
			channelRegister <- empty
		} else if err != nil {
			helper.PanicIfError(err)
		} else {
			channelRegister <- userExist
		}
	}()

	result := <-channelRegister

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
		result.IsLogin = domain.StatusOffline
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

		newUser, err := service.UserRepository.FindById(ctx, tx, result.ID)
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
	//tx := service.DB.WithContext(ctx).Begin()
	//helper.PanicIfError(tx.Error)
	//
	//defer helper.CommitOrRollback(tx)
	//
	//channelUser := make(chan domain.User, 1)
	//channelError := make(chan bool, 1)
	//
	//var wg sync.WaitGroup
	//
	//empty := domain.User{}
	//
	//// cek email exist
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	//	if err != nil {
	//		channelUser <- empty
	//	} else {
	//		channelUser <- user
	//	}
	//}()
	//
	//result := <-channelUser
	//
	//if result == empty {
	//	return api_response.AuthLoginResponse{}, errors.New("Email/Password salah.")
	//}
	//
	//// cek password
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
	//
	//	if err != nil {
	//		channelError <- true
	//	} else {
	//		channelError <- false
	//	}
	//}()
	//
	//yes := true
	//newError := <-channelError
	//if yes == newError {
	//	return api_response.AuthLoginResponse{}, errors.New("Email/Password salah.")
	//}
	//
	//// create token
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	id := strconv.Itoa(result.ID)
	//	token, err := helper.CreateToken(id)
	//	helper.PanicIfError(err)
	//	result.Token = sql.NullString{String: token, Valid: true}
	//
	//	channelUser <- result
	//}()
	//
	//result = <-channelUser
	//
	//// update user
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	user := service.UserRepository.Update(ctx, tx, result)
	//	channelUser <- user
	//}()
	//
	//result = <-channelUser
	//
	//// find user
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	user, err := service.UserRepository.FindById(ctx, tx, result.ID)
	//	helper.PanicIfError(err)
	//	channelUser <- user
	//
	//}()
	//
	//go func() {
	//	wg.Wait()
	//	close(channelUser)
	//	close(channelError)
	//}()
	//
	//responseUser := <-channelUser
	//
	//lembaga := domain.Lembaga{
	//	ID:            responseUser.Lembaga.ID,
	//	Nama:          responseUser.Lembaga.Nama,
	//	Alamat:        responseUser.Lembaga.Alamat,
	//	IdTipeLembaga: responseUser.Lembaga.IdTipeLembaga,
	//}
	//return api_response.AuthLoginResponse{
	//	Id_User: responseUser.ID,
	//	Nama:    responseUser.Nama,
	//	Email:   responseUser.Email,
	//	Token:   responseUser.Token,
	//	Lembaga: lembaga,
	//}, nil
	tx := service.DB.WithContext(ctx).Begin()
	helper.PanicIfError(tx.Error)

	defer helper.CommitOrRollback(tx)

	// Cek email exist
	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return api_response.AuthLoginResponse{}, errors.New("Email/Password salah.")
	}

	// Cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return api_response.AuthLoginResponse{}, errors.New("Email/Password salah.")
	}

	// Buat token
	id := strconv.Itoa(user.ID)
	token, err := helper.CreateToken(id)
	helper.PanicIfError(err)
	user.Token = sql.NullString{String: token, Valid: true}

	// Update user
	updatedUser := service.UserRepository.Update(ctx, tx, user)

	// Temukan ulang user
	foundUser, err := service.UserRepository.FindById(ctx, tx, updatedUser.ID)
	helper.PanicIfError(err)

	lembaga := domain.Lembaga{
		ID:            foundUser.Lembaga.ID,
		Nama:          foundUser.Lembaga.Nama,
		Alamat:        foundUser.Lembaga.Alamat,
		IdTipeLembaga: foundUser.Lembaga.IdTipeLembaga,
	}

	return api_response.AuthLoginResponse{
		Id_User: foundUser.ID,
		Nama:    foundUser.Nama,
		Email:   foundUser.Email,
		Token:   foundUser.Token,
		Lembaga: lembaga,
	}, nil
}
