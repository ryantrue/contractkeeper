package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ryantrue/contractkeeper/internal/config"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
