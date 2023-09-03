package routes

import (
	"github.com/julienschmidt/httprouter"
	"go_share/controller/auth_controller"
)

type AuthRoutesImpl struct {
	AuthController auth_controller.AuthController
}

func InitializeAuthComponents(authController auth_controller.AuthController) *AuthRoutesImpl {
	return &AuthRoutesImpl{
		AuthController: authController,
	}
}

func (auth *AuthRoutesImpl) AuthRoutesComponent(router *httprouter.Router) {
	router.POST("/api/v1/auth/register", auth.AuthController.Register)
	router.POST("/api/v1/auth/login", auth.AuthController.Login)
}
