package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-todo-apiservices"
	"fmt"
)

var DBConf services.ConfigurationDB

func GetDBConnection() (*sql.DB, error){
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		DBConf.User,
		DBConf.Password,
		DBConf.Server,
		DBConf.DataBase,
		)
	dbContext, err := sql.Open("postgres", connString)
	if(err != nil){
		return  dbContext, err
	}

	return dbContext, nil
}