// database/database.go
package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"auth-api/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto migrate model User
	DB.AutoMigrate(&models.User{})
}