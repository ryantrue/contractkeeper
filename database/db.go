package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ryantrue/contractkeeper/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")

	// Automatically migrate the schema
	err = DB.AutoMigrate(&models.Request{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
