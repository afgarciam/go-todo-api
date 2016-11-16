package routers

import (
	"github.com/gorilla/mux"
	"todoisAPI/controllers"
	"github.com/urfave/negroni"
)

func BuildUserRouter(router *mux.Router) (*mux.Router) {
	prefix := "/api/users"

	usr := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	usr.HandleFunc("/{id}",controllers.GetUserById).Methods("GET")

	router.PathPrefix(prefix).Handler(negroni.New(
		//negroni.HandlerFunc(services.VerifiyToken),
		negroni.Wrap(usr),
	))

	return router
}
