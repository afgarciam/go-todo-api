package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"todoisAPI/services"
	"fmt"
	"sync"
)


var once = sync.Once{}

func GetDBConnection() (*sql.DB, error){
	conf := &services.ConfigurationDB{}
	once.Do(func(){
		conf.Load()
	})
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		conf.User,
		conf.Password,
		conf.Server,
		conf.DataBase,
	)
	dbContext, err := sql.Open("postgres", connString)
	if(err != nil){
		return  dbContext, err
	}

	return dbContext, nil
}