package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"go_share/helper"
	"go_share/model/api/api_request"
	"go_share/model/api/api_response"
	"go_share/model/domain"
	"go_share/repository/user_repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
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

	empty := domain.User{}
	userExist, _ := service.UserRepository.FindByEmail(ctx, tx, request.Email)

	if userExist != empty {
		return api_response.AuthRegisterResponse{}, errors.New("Email telah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tiketStr := strconv.Itoa(int(time.Now().Unix()))
	tiket := sql.NullString{String: tiketStr, Valid: true}

	user := domain.User{
		Nama:      request.Nama,
		Email:     request.Email,
		Password:  string(hashedPassword),
		IdLembaga: request.IdLembaga,
		Tiket:     tiket,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	newUser, err := service.UserRepository.FindById(ctx, tx, user.IdUser)
	helper.PanicIfError(err)

	return api_response.AuthRegisterResponse{
		Email: newUser.Email,
		Tiket: tiketStr,
	}, nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request api_request.AuthLoginRequest, channelLogin chan<- interface{}) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// cek email exist
	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicIfError(err)

	// cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(err)
		fmt.Println("Password salah.")
	}

	// create token
	token, err := CreateToken(string(user.IdUser))
	user.Token = sql.NullString{String: token, Valid: true}

	// update user
	user = service.UserRepository.Update(ctx, tx, user)

	// find user
	user, err = service.UserRepository.FindById(ctx, tx, user.IdUser)

	response := api_response.AuthLoginResponse{
		Id_User: user.IdUser,
		Nama:    user.Nama,
		Email:   user.Email,
		Token:   token,
	}

	channelLogin <- response
}

func CreateToken(value string) (string, error) {
	claims := jwt.MapClaims{
		"id":  value,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Waktu kadaluwarsa token (contoh: 1 jam)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("key-token-rahasia") // Ganti dengan kunci rahasia yang kuat

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte("key-token-rahasia")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode tanda tangan tidak valid")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
