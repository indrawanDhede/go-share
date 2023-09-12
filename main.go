package main

import (
	"github.com/go-playground/validator"
	"go_share/app/database"
	"go_share/exception"
	"go_share/helper"
	"go_share/middleware"
	"go_share/routes/auth_routes"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.NewDBGorm()
	validate := validator.New()
	router := httprouter.New()

	authRoutes := auth_routes.InitializeAuthRoute(db, validate)
	authRoutes.AuthRoutesComponent(router)
	//userRoutes := user_routes.InitializeUserRoute(db, validate)
	//userRoutes.UserRoutesComponent(router)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router, db),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
