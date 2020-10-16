package main

import (
	"net/http"
	"fmt"
	"log"
	"go-todo-apirouters"
	"time"
	"go-todo-apiservices"
	"go-todo-apidao"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func(){
		dao.DBConf =  services.LoadDataBaseConfig()
	})

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
