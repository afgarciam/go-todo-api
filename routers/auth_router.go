package routers

import (
	"github.com/gorilla/mux"
	"todoisAPI/controllers"
)

func BuildAuthRouter(router *mux.Router)  (*mux.Router){
	prefix := "/api/auth"

	auth := router.PathPrefix(prefix).Subrouter().StrictSlash(true)
	auth.HandleFunc("/login", controllers.Login).Methods("POST")

	return router
}
