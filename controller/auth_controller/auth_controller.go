package auth_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
