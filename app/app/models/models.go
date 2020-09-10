package models

// Profile struct is for general user profile
type Profile struct {
	Name     string
	Email    string
	MobileNo uint
	Password string
}

// User model is for customers objects
type User struct {
	ID      uint
	Orders  []Order
	Profile Profile `gorm:"embedded"`
}

// DeliveryPerson model is for Delivery Person objects
type DeliveryPerson struct {
	ID            uint
	Profile       Profile `gorm:"embedded"`
	CurrentOrders []Order
	CurrentArea   uint
	CurrentLoc    uint
}

// Cookie model is for Cookie Objects
type Cookie struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	Quantity    uint
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
	DeliveryLocation  uint
}
