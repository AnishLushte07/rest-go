package db

import (
	"log"
	"rest-go/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB
var err error

func Connect(dbName string) {
	DbInstance, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	DbInstance.AutoMigrate(&entities.Product{})
	log.Println("DB migration completed")
}
