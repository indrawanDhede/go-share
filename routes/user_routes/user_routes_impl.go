package user_routes

import (
	"github.com/julienschmidt/httprouter"
	"go_share/controller/user_controller"
)

type UserRoutesImpl struct {
	UserController user_controller.UserController
}

func NewUserRoutes(userController user_controller.UserController) *UserRoutesImpl {
	return &UserRoutesImpl{UserController: userController}
}

func (route UserRoutesImpl) UserRoutesComponent(router *httprouter.Router) {
	router.GET("/api/v1/user", route.UserController.FindAll)
}
