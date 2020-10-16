package dao

import (
	"database/sql"
	"fmt"
	"go-todo-api/services"

	_ "github.com/lib/pq"
)

var DBConf services.ConfigurationDB

func GetDBConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		DBConf.User,
		DBConf.Password,
		DBConf.Server,
		DBConf.DataBase,
	)
	dbContext, err := sql.Open("postgres", connString)
	if err != nil {
		return dbContext, err
	}

	return dbContext, nil
}
