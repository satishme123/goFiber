package database

import (
	"log"
	"os"
	models "simple-api-with-go-fiber/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstances struct {
	Db *gorm.DB
}

var Database DbInstances

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Databse open failed", err)
		os.Exit(1)
	}
	log.Println("connected to database successfully")
	log.Println()
	db.AutoMigrate(&models.User{}, &models.Product{})
	Database = DbInstances{Db: db}
}
