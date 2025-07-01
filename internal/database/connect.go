package database

import (
	model "example/todo-api/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(pathToDB string) {

	var err error

	database, err := gorm.Open(sqlite.Open(pathToDB), &gorm.Config{})
	if err != nil {
		panic("failed connect to DB")
	}

	err = database.AutoMigrate(&model.TodoModel{}, &model.UserModel{})
	if err != nil {
		log.Fatal("Error on MIGRATE DB")
	}

	DB = database
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := GetDB().DB()
	if err != nil {
		log.Fatal("Error on GET DB", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Error on CLOSE DB", err)
	}
}
