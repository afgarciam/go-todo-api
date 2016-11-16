package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"todoisAPI/services"
	"fmt"
)

func GetDBConnection() (*sql.DB, error){
	conf := &services.ConfigurationDB{}
	conf.Load()
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