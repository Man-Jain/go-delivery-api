package services

import (
	"app/app/models"
)

// QueryOrders get all current orders
func QueryOrders() (*[]models.Order, error) {
	orders := []models.Order{}
	result := DB.Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return &orders, nil
}

// QueryOrdersUser get all current orders of a user
func QueryOrdersUser(userid uint) (*[]models.Order, error) {
	orders := []models.Order{}
	result := DB.Where("user_id = ?", userid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return &orders, nil
}

// QueryOrdersDeliveryAgent get all current orders of delivery agent
func QueryOrdersDeliveryAgent(deliverypersonid uint) (*[]models.Order, error) {
	orders := []models.Order{}
	result := DB.Where("delivery_person_id = ?", deliverypersonid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return &orders, nil
}

// InsertOrder create a new order
func InsertOrder(jsonData map[string]interface{}) (*models.Order, error) {
	var order, sameAreaOrder models.Order
	var deliveryPerson models.DeliveryPerson

	deliveryArea := uint(jsonData["delivery_area"].(float64))
	if sameResult := DB.Where("delivery_area = ?", deliveryArea).First(&sameAreaOrder); sameResult.Error == nil {
		deliverPerson := models.DeliveryPerson{}
		deliveryPersonID := sameAreaOrder.DeliveryPersonID
		if result := DB.First(&deliveryPerson, deliveryPersonID); result.Error == nil {
			println("deliveryArea", deliveryArea)
			println("deliverPerson", deliverPerson.CurrentLoc)
			estimatedDelivery := deliveryArea - deliverPerson.CurrentLoc
			order.CookieID = uint(jsonData["cookie_id"].(float64))
			order.Quantity = uint(jsonData["quantity"].(float64))
			order.UserID = uint(jsonData["user_id"].(float64))
			order.DeliveryPersonID = deliveryPersonID
			order.EstimatedDelivery = estimatedDelivery
			order.Status = "EnRoute"
			order.DeliveryArea = deliveryArea

			if result := DB.Create(&order); result.Error != nil {
				return &order, result.Error
			}
		}
	} else {
		var deliveryPersonID uint = 1
		var estimatedDelivery uint = 32

		order.CookieID = uint(jsonData["cookie_id"].(float64))
		order.Quantity = uint(jsonData["quantity"].(float64))
		order.UserID = uint(jsonData["user_id"].(float64))
		order.DeliveryPersonID = deliveryPersonID
		order.EstimatedDelivery = estimatedDelivery
		order.Status = "EnRoute"
		order.DeliveryArea = deliveryArea

		if result := DB.Create(&order); result.Error != nil {
			return &order, result.Error
		}
	}

	return &order, nil
}

// QueryOrder returns a single order object
func QueryOrder(id int) (*models.Order, error) {
	order := models.Order{}
	result := DB.First(&order, id)
	println(result)
	return &order, nil
}
