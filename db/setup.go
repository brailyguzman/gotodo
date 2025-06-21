package db

import (
	"log"

	"github.com/brailyguzman/gotodo/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Database: %v", err)
	}

	log.Println("Database connection established successfully")
}

func Migrate() {
	// check if tables already exist
	if !DB.Migrator().HasTable("users") {
		err := DB.AutoMigrate(&models.User{})

		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	log.Println("Database migration completed successfully")
}
