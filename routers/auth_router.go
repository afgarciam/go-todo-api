package routers

import (
	"github.com/gorilla/mux"
	"go-todo-apicontrollers"
)

func BuildAuthRouter(router *mux.Router)  (*mux.Router){
	var authCtrl *controllers.AuthController
	prefix := "/api/auth"

	auth := router.PathPrefix(prefix).Subrouter().StrictSlash(true)
	auth.HandleFunc("/login", authCtrl.Login).Methods("POST")

	return router
}
