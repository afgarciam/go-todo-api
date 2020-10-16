package routers

import (
	"go-todo-api/controllers"
	services "go-todo-api/services"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func BuildUserRouter(router *mux.Router) *mux.Router {
	var userCtrl *controllers.UserController
	prefix := "/api/users"

	usr := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	usr.HandleFunc("/{id}", userCtrl.GetById).Methods("GET")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(services.VerifyToken),
		negroni.Wrap(usr),
	))

	return router
}
