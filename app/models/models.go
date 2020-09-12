package models

import "github.com/revel/revel"

// Profile struct is for general user profile
type Profile struct {
	Name     string
	Email    string `gorm:"unique"`
	MobileNo uint   // `gorm:"unique"`
	Password []byte
	Address  uint
}

// User model is for customers objects
type User struct {
	ID      uint
	Orders  []Order
	Profile Profile `gorm:"embedded"`
	Root    bool
}

// Validate user Model
func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Profile.Email)
	v.Required(user.Profile.Password)
	v.MinSize(user.Profile.Password, 6)
	v.Email(user.Profile.Email)
}

// DeliveryPerson model is for Delivery Person objects
type DeliveryPerson struct {
	ID            uint
	Profile       Profile `gorm:"embedded"`
	CurrentOrders []Order
	// OrderArea     uint
	CurrentLoc uint
	OnDelivery bool
}

// Validate user Model
func (dp *DeliveryPerson) Validate(v *revel.Validation) {
	v.Required(dp.Profile.Email)
	v.Required(dp.Profile.Password)
	v.Email(dp.Profile.Email)
}

// Cookie model is for Cookie Objects
type Cookie struct {
	ID                uint
	Name, Description string
	Price             uint
	Quantity          uint
	Orders            []Order
}

// Validate user Model
func (cookie *Cookie) Validate(v *revel.Validation) {
	v.Required(cookie.Name)
	v.Required(cookie.Description)
	// v.Min(int(cookie.Price), 1)
	// v.Min(int(cookie.Quantity), 1)
}

// Order model is for Order Objects
type Order struct {
	ID                uint
	CookieID          uint
	Quantity          uint
	UserID            uint
	DeliveryPersonID  uint
	EstimatedDelivery uint
	Status            string
	DeliveryArea      uint
}
