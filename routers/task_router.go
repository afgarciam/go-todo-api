package routers

import (
	"github.com/gorilla/mux"
	"todoisAPI/controllers"
	"github.com/urfave/negroni"
	"todoisAPI/services"
)

func BuildTaskRouter(router *mux.Router) (*mux.Router) {
	prefix := "/api/tasks"

	task := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	task.HandleFunc("/",controllers.GetAllTask).Methods("GET")
	task.HandleFunc("/",controllers.CreateTask).Methods("POST")
	task.HandleFunc("/{id}",controllers.UpdateTask).Methods("PUT")
	task.HandleFunc("/{id}",controllers.DeleteTask).Methods("DELETE")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(services.VerifiyToken),
		negroni.Wrap(task),
	))

	return router
}

