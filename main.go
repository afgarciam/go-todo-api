package main

import (
	"net/http"
	"fmt"
	"log"
	"todoisAPI/services"
	"todoisAPI/routers"
	"time"
)

func main() {
	conf := services.ConfigurationDB{}
	err := conf.Load()
	if(err != nil){
		log.Fatal(err)
	}

	s:= &http.Server{
		Addr: ":9003",
		Handler: routers.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server app in http://127.0.0.1:9003")
	log.Fatal(s.ListenAndServe())
}



