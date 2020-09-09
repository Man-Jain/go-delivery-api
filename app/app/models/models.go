package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	Name     string
	Email    string
	MobileNo int
}

type User struct {
	gorm.Model
	Orders  []Order
	Profile Profile `gorm:"embedded"`
}

type DeliveryPerson struct {
	gorm.Model
	Profile       Profile `gorm:"embedded"`
	CurrentOrders []Order
	CurrentArea   string
	CurrentLoc    string
}

type Cookie struct {
	gorm.Model
	Name        string
	Description string
	Price       int
}

type Order struct {
	gorm.Model
	Cookies           []Cookie `gorm:"many2many:cookie_orders"`
	UserID            int
	DeliveryPersonID  int
	EstimatedDelivery int
}
