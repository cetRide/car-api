package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)

var database *gorm.DB
func InitDB(dbUrI string) {
	connection, err := gorm.Open("postgres", dbUrI)
	if err != nil {
		panic(err)
	}
	database = connection
	MigrateDatabase(database)
}

func GetDB() *gorm.DB {
	return database
}
