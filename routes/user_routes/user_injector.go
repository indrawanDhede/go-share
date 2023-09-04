//go:build wireinject
// +build wireinject

package user_routes

import (
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"go_share/controller/user_controller"
	"go_share/repository/user_repository"
	"go_share/service/user_service"
)

func InitializeUserRoute(db *sql.DB, validate *validator.Validate) *UserRoutesImpl {
	wire.Build(
		user_repository.NewUserRepository, user_service.NewUserService, user_controller.NewUserController, NewUserRoutes,
	)
	return nil
}
