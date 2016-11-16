package routers

import (
	"github.com/gorilla/mux"
	"todoisAPI/controllers"
	"github.com/urfave/negroni"
	"todoisAPI/services"
)

func BuildUserRouter(router *mux.Router) (*mux.Router) {
	var userCtrl *controllers.UserController
	prefix := "/api/users"

	usr := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	usr.HandleFunc("/{id}",userCtrl.GetById).Methods("GET")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(services.VerifyToken),
		negroni.Wrap(usr),
	))

	return router
}
