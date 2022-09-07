package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var dbInstance *gorm.DB = nil

func Connect() *gorm.DB {
	if dbInstance == nil {
		dbInstance, err := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to db")
		}
		return dbInstance
	} else {
		return dbInstance
	}
}
