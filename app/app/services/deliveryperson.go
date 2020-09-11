package services

import (
	"app/app/models"
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
func InsertDeliveryPerson(obj models.DeliveryPerson) (*models.DeliveryPerson, error) {

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
