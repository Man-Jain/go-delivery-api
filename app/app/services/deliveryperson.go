package services

import (
	"app/app/models"

	"golang.org/x/crypto/bcrypt"
)

// QueryAllDeliveryPeople return all users of application
func QueryAllDeliveryPeople() (*[]models.DeliveryPerson, error) {
	deliveryPeople := []models.DeliveryPerson{}
	result := DB.Find(&deliveryPeople)
	println(result)
	println(result.RowsAffected)
	return &deliveryPeople, nil
}

// QueryDeliveryPerson return a single user object
func QueryDeliveryPerson(id uint) (*models.DeliveryPerson, error) {
	dp := models.DeliveryPerson{}
	result := DB.First(&dp, id)
	println(result.RowsAffected)
	return &dp, nil
}

// InsertDeliveryPerson will insert a delivery person in db
func InsertDeliveryPerson(jsonData map[string]interface{}) (*models.DeliveryPerson, error) {
	password := jsonData["password"].(string)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	obj := models.DeliveryPerson{
		Profile: models.Profile{
			Email:    jsonData["email"].(string),
			Password: hashedPassword,
		},
	}
	if result := DB.Create(&obj); result.Error != nil {
		return &obj, result.Error
	}

	return &obj, nil
}

// UpdateDelivery return a single user object
func UpdateDelivery(id int) (*models.Order, error) {
	order := models.Order{}
	result := DB.First(&order, id)
	order.Status = "Delivered"

	DB.Save(&order)
	println(result.RowsAffected)
	return &order, nil
}
