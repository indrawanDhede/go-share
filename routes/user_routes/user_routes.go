package user_routes

import "github.com/julienschmidt/httprouter"

type UserRoutes interface {
	UserRoutesComponent(router *httprouter.Router)
}
