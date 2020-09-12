package services

import (
	"app/app/models"
	"fmt"
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
		dbHost := os.Getenv("dbHost")
		username := os.Getenv("username")
		dbName := os.Getenv("dbName")
		password := os.Getenv("password")

		dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

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
