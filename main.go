package main

import (
	"go_share/app"
	"go_share/controller/auth_controller"
	"go_share/exception"
	"go_share/helper"
	"go_share/repository/user_repository"
	"go_share/routes"
	"go_share/service/auth_service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()
	authRepository := user_repository.NewUserRepository()
	authService := auth_service.NewAuthService(authRepository, db, validate)
	authController := auth_controller.NewAuthController(authService)

	authRoutes := routes.InitializeAuthComponents(authController)
	authRoutes.AuthRoutesComponent(router)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
