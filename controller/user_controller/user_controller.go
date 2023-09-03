package user_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
