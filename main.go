package main

import (
	"fmt"
	"go-todo-api/dao"
	"go-todo-api/routers"
	"go-todo-api/services"
	"log"
	"net/http"
	"sync"
	"time"
)

var once sync.Once

func main() {
	once.Do(func() {
		dao.DBConf = services.LoadDataBaseConfig()
	})

	s := &http.Server{
		Addr:           ":9003",
		Handler:        routers.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server app in http://127.0.0.1:9003")
	log.Fatal(s.ListenAndServe())
}
