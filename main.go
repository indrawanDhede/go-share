package main

import (
	"go_share/app"
	"go_share/controller/auth_controller"
	"go_share/exception"
	"go_share/helper"
	"go_share/repository/user_repository"
	"go_share/service/auth_service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	authRepository := user_repository.NewUserRepository()
	authService := auth_service.NewAuthService(authRepository, db, validate)
	authController := auth_controller.NewAuthController(authService)

	router := httprouter.New()

	router.POST("/api/v1/auth/register", authController.Register)
	router.POST("/api/v1/auth/login", authController.Login)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
