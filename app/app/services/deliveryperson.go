package services

import (
	"app/app/models"

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
		return &order, result.Error
	}
	order.Status = "Delivered"

	deliveryPerson := models.DeliveryPerson{}
	if result := DB.First(&deliveryPerson, order.DeliveryPersonID); result.Error == nil {
		return &order, result.Error
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Model(&deliveryPerson).Update("on_delivery", "1").Where("id", order.DeliveryPersonID); result.Error != nil {
			println("This is an error")
			return result.Error
		}

		if result := tx.Model(&deliveryPerson).Update("current_loc", order.DeliveryArea).Where("id", order.DeliveryPersonID); result.Error != nil {
			println("This is an error")
			return result.Error
		}

		if result := tx.Model(&order).Update("status", "Delivered").Where("id", order.ID); result.Error != nil {
			println("This is an error")
			return result.Error
		}

		return nil
	})

	return &order, err
}
