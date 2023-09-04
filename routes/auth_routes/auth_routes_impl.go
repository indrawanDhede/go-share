package auth_routes

import (
	"github.com/julienschmidt/httprouter"
	"go_share/controller/auth_controller"
)

type AuthRoutesImpl struct {
	AuthController auth_controller.AuthController
}

func NewAuthRoutes(authController auth_controller.AuthController) *AuthRoutesImpl {
	return &AuthRoutesImpl{
		AuthController: authController,
	}
}

func (route *AuthRoutesImpl) AuthRoutesComponent(router *httprouter.Router) {
	router.POST("/api/v1/auth/register", route.AuthController.Register)
	router.POST("/api/v1/auth/login", route.AuthController.Login)
}
