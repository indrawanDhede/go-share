package routes

import "github.com/julienschmidt/httprouter"

type AuthRoutes interface {
	AuthRoutesComponent(router *httprouter.Router)
}
