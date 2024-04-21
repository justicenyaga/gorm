package main

import (
	"encoding/json"
	"fmt"
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

	// dbtools.CreateTable(&models.Student{})

	student := models.Student{
		ID:   4,
		Name: "Robert Downey",
		Age:  44,
	}
	dbtools.Save(&student)
	fmt.Println("Student added to the database")
}
