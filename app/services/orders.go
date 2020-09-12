package services

import (
	"app/app/models"
	"errors"

	"gorm.io/gorm"
)

// AbsIntUtil is absolute for integer
func AbsIntUtil(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

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
func InsertOrder(userID uint, cookieID uint, quantity uint) (*models.Order, error) {
	var order, sameAreaOrder models.Order
	var deliveryPerson models.DeliveryPerson
	var customer models.User
	var cookie models.Cookie

	if result := DB.First(&customer, userID); result.Error != nil {
		return &order, result.Error
	}

	deliveryArea := customer.Profile.Address

	if result := DB.First(&cookie, cookieID); result.Error != nil {
		return &order, result.Error
	}

	if cookie.Quantity < quantity {
		println("This is negative")
		return &order, errors.New("Not enough Quantity")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {

		cookie.Quantity = cookie.Quantity - quantity
		if result := tx.Save(&cookie); result.Error != nil {
			println("This is an error 1")
			return result.Error
		}

		println("deliveryArea", deliveryArea)

		if result := tx.Where("delivery_	area = ?", deliveryArea).First(&sameAreaOrder); result.Error == nil {
			println("FOund delivery in same area")
			deliveryPersonID := sameAreaOrder.DeliveryPersonID
			println("deliveryPersonID", deliveryPersonID)
			if result := tx.First(&deliveryPerson, deliveryPersonID); result.Error != nil {
				println("This is an error 2", result.Error)
				return result.Error
			}

			deliveryPerson.OnDelivery = true
			if result := tx.Model(&deliveryPerson).Update("on_delivery", "1").Where("id", deliveryPersonID); result.Error != nil {
				println("This is an error 3")
				return result.Error
			}

			estimatedDelivery := int(deliveryArea - deliveryPerson.CurrentLoc)
			if estimatedDelivery < 0 {
				estimatedDelivery = -estimatedDelivery
			}

			order.CookieID = cookieID
			order.Quantity = quantity
			order.UserID = userID
			order.DeliveryPersonID = deliveryPersonID
			order.EstimatedDelivery = uint(estimatedDelivery)
			order.Status = "EnRoute"
			order.DeliveryArea = deliveryArea

			if result := tx.Create(&order); result.Error != nil {
				println("This is an error 4")
				return result.Error
			}

		} else {
			println("FOund not delivery in same area")

			var minEstDelivery int = 99999
			var deliveryPersonID uint = 0
			println("min", minEstDelivery)
			deliveryPeople, _ := QueryAllDeliveryPeople()

			for _, value := range *(deliveryPeople) {
				currentLoc := value.CurrentLoc
				estDelivery := AbsIntUtil(int(currentLoc) - int(deliveryArea))
				println("minimunnn", currentLoc, deliveryArea, minEstDelivery, value.ID)
				if estDelivery < minEstDelivery {
					minEstDelivery = estDelivery
					deliveryPersonID = value.ID
				}
			}

			println("Minimu is ", minEstDelivery, deliveryPersonID)

			if result := DB.First(&deliveryPerson, deliveryPersonID); result.Error != nil {
				println("This is an error 22", result.Error)
				return result.Error
			}

			order.CookieID = cookieID
			order.Quantity = quantity
			order.UserID = userID
			order.DeliveryPersonID = deliveryPersonID
			order.EstimatedDelivery = uint(minEstDelivery)
			order.Status = "EnRoute"
			order.DeliveryArea = deliveryArea

			if result := tx.First(&deliveryPerson, deliveryPersonID); result.Error != nil {
				println("This is an error")
				return result.Error
			}

			deliveryPerson.OnDelivery = true
			if result := tx.Model(&deliveryPerson).Update("on_delivery", "1").Where("id", deliveryPersonID); result.Error != nil {
				println("This is an error")
				return result.Error
			}

			if result := tx.Create(&order); result.Error != nil {
				println("This is an error")
				return result.Error
			}
		}

		return nil
	})

	return &order, err
}

// QueryOrder returns a single order object
func QueryOrder(id int) (*models.Order, error) {
	order := models.Order{}
	result := DB.First(&order, id)
	println(result)
	return &order, nil
}
