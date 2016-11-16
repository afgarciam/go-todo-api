package routers

import (
	"github.com/gorilla/mux"
	"todoisAPI/controllers"
	"github.com/urfave/negroni"
	"todoisAPI/services"
)

func BuildTaskRouter(router *mux.Router) (*mux.Router) {
	var taskCtrl *controllers.TaskController
	prefix := "/api/tasks"

	task := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	task.HandleFunc("/",taskCtrl.GetAll).Methods("GET")
	task.HandleFunc("/",taskCtrl.Create).Methods("POST")
	task.HandleFunc("/{id}",taskCtrl.Update).Methods("PUT")
	task.HandleFunc("/{id}",taskCtrl.Delete).Methods("DELETE")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(services.VerifyToken),
		negroni.Wrap(task),
	))

	return router
}

