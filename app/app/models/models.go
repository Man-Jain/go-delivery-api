package models

type Profile struct {
	Name     string
	Email    string
	MobileNo int
	Password string
}

type User struct {
	ID      uint
	Orders  []Order
	Profile Profile `gorm:"embedded"`
}

type DeliveryPerson struct {
	ID            uint
	Profile       Profile `gorm:"embedded"`
	CurrentOrders []Order
	CurrentArea   string
	CurrentLoc    string
}

type Cookie struct {
	ID          uint
	Name        string
	Description string
	Price       int
}

type Order struct {
	ID                uint
	Cookies           []Cookie `gorm:"many2many:cookie_orders"`
	UserID            int
	DeliveryPersonID  int
	EstimatedDelivery int
}
