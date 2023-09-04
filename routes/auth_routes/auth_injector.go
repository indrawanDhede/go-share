//go:build wireinject
// +build wireinject

package auth_routes

import (
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"go_share/controller/auth_controller"
	"go_share/repository/user_repository"
	"go_share/service/auth_service"
)

func InitializeAuthRoute(db *sql.DB, validate *validator.Validate) *AuthRoutesImpl {
	wire.Build(
		user_repository.NewUserRepository, auth_service.NewAuthService, auth_controller.NewAuthController, NewAuthRoutes,
	)
	return nil
}
