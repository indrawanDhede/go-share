package auth_controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go_share/helper"
	"go_share/model/api"
	"go_share/model/api/api_request"
	"go_share/service/auth_service"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService auth_service.AuthService
}

func NewAuthController(authService auth_service.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	authRegisterRequest := api_request.AuthRegisterRequest{}
	err := decoder.Decode(&authRegisterRequest)
	helper.PanicIfError(err)

	authResponse, err := controller.AuthService.Register(request.Context(), authRegisterRequest)

	if err != nil {
		response := api.ApiResponseGeneral{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}

		helper.WriteToResponse(writer, response)
	} else {
		response := api.ApiResponseGeneral{
			Code:   200,
			Status: "OK",
			Data:   authResponse,
		}

		helper.WriteToResponse(writer, response)
	}
}

func (controller AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	authLoginRequest := api_request.AuthLoginRequest{}
	err := decoder.Decode(&authLoginRequest)
	helper.PanicIfError(err)

	channelLogin := make(chan interface{})
	defer close(channelLogin)

	go controller.AuthService.Login(request.Context(), authLoginRequest, channelLogin)

	data := <-channelLogin

	fmt.Println("done..")

	response := api.ApiResponseGeneral{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	helper.WriteToResponse(writer, response)
}
