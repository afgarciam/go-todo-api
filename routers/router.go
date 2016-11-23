package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
)


func Router() (*negroni.Negroni) {
	c := cors.Default()

	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.Use(c)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GOLANG API Restful");
	})
	router = BuildUserRouter(router)
	router = BuildAuthRouter(router)
	router = BuildTaskRouter(router)

	n.UseHandler(router)
	return  n
}