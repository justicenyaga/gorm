package dbtools

import (
	"fmt"
	"gorm-mysql/models"
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

func CreateTable(object ...interface{}) {
	db := connect()

	err := db.AutoMigrate(object...)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Table(s) created successfully")
}

func Save(object interface{}) {
	db := connect()

	db.Create(object)
}

func Select(students []models.Student) {
	db := connect()

	db.Find(&students)

	for _, student := range students {
		fmt.Printf("%d\t%s\t%d\n", student.ID, student.Name, student.Age)
	}
}

func SingleUpdate(object interface{}, data map[string]interface{}) int64 {
	db := connect()

	result := db.Model(object).Updates(data)

	return result.RowsAffected
}

func MultipleUpdate(object interface{}, whereClause string, data map[string]interface{}) int64 {
	db := connect()

	result := db.Model(object).Where(whereClause).Updates(data)

	return result.RowsAffected
}

func SingleDelete(object interface{}) int64 {
	db := connect()

	result := db.Delete(object)

	return result.RowsAffected
}

func MultipleDelete(object interface{}, whereClause string) int64 {
	db := connect()

	result := db.Where(whereClause).Delete(object)

	return result.RowsAffected
}
