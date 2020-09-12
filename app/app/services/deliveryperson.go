package services

import (
	"app/app/models"
	"errors"

	"gorm.io/gorm"
)

// QueryAllDeliveryPeople return all users of application
func QueryAllDeliveryPeople() (*[]models.DeliveryPerson, error) {
	deliveryPeople := []models.DeliveryPerson{}
	result := DB.Preload("CurrentOrders").Find(&deliveryPeople)
	println(deliveryPeople)
	println(result.RowsAffected)
	return &deliveryPeople, nil
}

// QueryDeliveryPerson return a single user object
func QueryDeliveryPerson(id uint) (*models.DeliveryPerson, error) {
	dp := models.DeliveryPerson{}
	result := DB.Preload("CurrentOrders").First(&dp, id)
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
	if result := DB.First(&order, id); result.Error != nil {
		println("This is the error 0")
		return &order, result.Error
	}
	if order.Status == "Delivered" {
		println("Already Delivered")
		return &order, errors.New("Already Delivered")
	}
	println(order.ID, order.Status)
	order.Status = "Delivered"
	println(order.ID, order.Status)

	deliveryPerson := models.DeliveryPerson{}
	deliveryPersonID := order.DeliveryPersonID
	if result := DB.First(&deliveryPerson, deliveryPersonID); result.Error != nil {
		println("Error with Dleivery Guy")
		return &order, result.Error
	}

	deliveryPerson.OnDelivery = false
	deliveryPerson.CurrentLoc = order.DeliveryArea
	println(deliveryPerson.Profile.Email, deliveryPersonID)

	err := DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&order); result.Error != nil {
			println("This is an error 3")
			return result.Error
		}

		if result := tx.Save(&deliveryPerson); result.Error != nil {
			println("This is an error 1")
			return result.Error
		}
		return nil
	})

	return &order, err
}
