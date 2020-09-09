package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Email  string
	Orders []Order
	Type   string
}

type Cookie struct {
	gorm.Model
	Name        string
	Description string
	Price       int
}

type Order struct {
	gorm.Model
	Cookies           []Cookie
	Buyer             User
	DeliveryAgent     User
	EstimatedDelivery int
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func main() {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

}
