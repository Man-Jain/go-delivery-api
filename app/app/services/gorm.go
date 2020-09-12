package services

import (
	"app/app/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB instance connects to the Database
var DB *gorm.DB

// InitDB initialises DB instance
func InitDB() {
	var err error
	prod := os.Getenv("PROD")
	if prod == "True" {
		println("Using Postgres")
		dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		println("Using Sqlite")
		DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.DeliveryPerson{}, &models.Cookie{}, &models.Order{})
}
