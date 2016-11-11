package services

import (
	"os"
	"encoding/json"
)

type ConfigurationDB struct {
	DataBase string
	Server   string
	Port     string
	User     string
	Password string
}

func (c *ConfigurationDB)  Load() (error) {
	file, err := os.Open("./configurations/db_config.json")
	if (err != nil) {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if (err != nil) {
		return err
	}
	return nil
}