package models

// Profile struct is for general user profile
type Profile struct {
	Name     string
	Email    string
	MobileNo uint
	Password []byte
}

// User model is for customers objects
type User struct {
	ID      uint
	Orders  []Order
	Profile Profile `gorm:"embedded"`
	Root    bool
}

// DeliveryPerson model is for Delivery Person objects
type DeliveryPerson struct {
	ID            uint
	Profile       Profile `gorm:"embedded"`
	CurrentOrders []Order
	// OrderArea     uint
	CurrentLoc uint
}

// Cookie model is for Cookie Objects
type Cookie struct {
	ID                uint
	Name, Description string
	Price             uint
	Quantity          uint
	Orders            []Order
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
