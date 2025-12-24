package db

import (
	"log"
	"time"

	"hla_finder/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect_MySql_Server() {
	// Fixed for Docker – uses the MySQL service name
	dsn := "root:root@tcp(mysql:3306)/hla_db2?charset=utf8mb4&parseTime=true&loc=Local"

	log.Println("Attempting to connect to MySQL (host: mysql)...")

	var err error
	for attempt := 1; attempt <= 20; attempt++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Extra safety: test the connection
			sqlDB, _ := DB.DB()
			if sqlDB.Ping() == nil {
				log.Println("MySQL connected successfully!")
				return
			}
		}

		log.Printf("MySQL not ready yet (attempt %d/20): %v", attempt, err)
		time.Sleep(2 * time.Second) // wait 2 seconds before next try
	}

	// If all attempts fail
	panic("Failed to connect to MySQL after 20 attempts: " + err.Error())
}

func Create_Schema() {
	if DB == nil {
		log.Fatal("Database connection is nil. Connect first.")
	}

	err := DB.AutoMigrate(&models.User{}, &models.HLA{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	log.Println("Database schema migrated successfully (User, HLA tables created/updated)")
}