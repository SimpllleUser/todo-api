package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type TodoModel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null"`
	Description string `gorm:"null"`
	Completed   bool   `gorm:"default:false"`
}

const pathToDB = "./todos-db.db"

func ConnectDatabase() {

	var err error

	database, err := gorm.Open(sqlite.Open(pathToDB), &gorm.Config{})
	if err != nil {
		panic("failed connect to DB")
	}

	err = database.AutoMigrate(&TodoModel{})
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
