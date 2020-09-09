package controllers

import (
	"app/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// dsn := "user=postgres password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.DeliveryPerson{}, &models.Cookie{}, &models.Order{})
}
