package main

import (
	"encoding/json"
	"gorm-mysql/dbtools"
	"gorm-mysql/models"
	"log"
	"os"
)

type Configuration struct {
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var conf Configuration

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbtools.DBInitializer(conf.DataSourceName)

	dbtools.CreateTable(&models.Student{})
}
