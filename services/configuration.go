package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigurationDB struct {
	DataBase string
	Server   string
	Port     string
	User     string
	Password string
}


func  LoadDataBaseConfig() (ConfigurationDB) {
	var conf ConfigurationDB

	file, err := ioutil.ReadFile("./configurations/db_config.json")
	if (err != nil) {
		log.Fatal(err)
	}
	err = json.Unmarshal(file,&conf)
	if (err != nil) {
		log.Fatal(err)
	}
	return conf
}