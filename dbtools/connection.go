package dbtools

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dataSourceName string

func DBInitializer(dsn string) {
	dataSourceName = dsn
}

func connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
