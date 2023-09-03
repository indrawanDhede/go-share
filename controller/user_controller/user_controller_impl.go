package user_controller

import (
	"github.com/julienschmidt/httprouter"
	"go_share/helper"
	"go_share/model/api"
	"go_share/service/user_service"
	"net/http"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userResponse := controller.UserService.FindAll(request.Context())
	response := api.ApiResponseGeneral{
		Total:  len(userResponse),
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponse(writer, response)
}
