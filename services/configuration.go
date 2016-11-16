package services

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigurationDB struct {
	DataBase string
	Server   string
	Port     string
	User     string
	Password string
}

func (c *ConfigurationDB)  Load() (error) {
	file, err := ioutil.ReadFile("./configurations/db_config.json")
	if (err != nil) {
		return err
	}
	err = json.Unmarshal(file,c)
	if (err != nil) {
		return err
	}
	return nil
}