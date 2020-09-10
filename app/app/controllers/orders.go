package controllers

import (
	"app/app/models"
	"net/http"

	"github.com/revel/revel"
)

// Orders Controller with embedded App
type Orders struct {
	App
}

// GetOrders get all current orders
func (c *Orders) GetOrders() revel.Result {
	orders := []models.Order{}
	result := DB.Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return c.RenderJSON(orders)
}

// GetOrdersUser get all current orders of a user
func (c *Orders) GetOrdersUser(userid uint) revel.Result {
	orders := []models.Order{}
	result := DB.Where("user_id = ?", userid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return c.RenderJSON(orders)
}

// GetOrdersDeliveryAgent get all current orders of delivery agent
func (c *Orders) GetOrdersDeliveryAgent(deliverypersonid uint) revel.Result {
	orders := []models.Order{}
	result := DB.Where("delivery_person_id = ?", deliverypersonid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return c.RenderJSON(orders)
}

// CreateOrder create a new order
func (c *Orders) CreateOrder() revel.Result {
	var order, sameAreaOrder models.Order
	var jsonData map[string]interface{}
	var deliveryPerson models.DeliveryPerson

	c.Params.BindJSON(&jsonData)

	deliveryArea := jsonData["delivery_area"].(uint)
	if sameResult := DB.Where("estimated_delivery = ?", deliveryArea).First(&sameAreaOrder); sameResult.Error == nil {
		deliverPerson := models.DeliveryPerson{}
		deliveryPersonID := sameAreaOrder.DeliveryPersonID
		if result := DB.First(&deliveryPerson, deliveryPersonID); result.Error == nil {
			estimatedDelivery := deliveryArea - deliverPerson.CurrentLoc
			order.CookieID = jsonData["cookie_id"].(uint)
			order.Quantity = jsonData["quantity"].(uint)
			order.UserID = jsonData["user_id"].(uint)
			order.DeliveryPersonID = deliveryPersonID
			order.EstimatedDelivery = estimatedDelivery
			order.Status = "En Route"
			order.DeliveryArea = deliveryArea

			if result := DB.Create(&order); result.Error != nil {
				c.Response.Status = http.StatusBadRequest
				return c.RenderJSON("Cannot Create Order")
			}
		}
	} else {
		var deliveryPersonID uint = 1
		var estimatedDelivery uint = 32

		order.CookieID = jsonData["cookie_id"].(uint)
		order.Quantity = jsonData["quantity"].(uint)
		order.UserID = jsonData["user_id"].(uint)
		order.DeliveryPersonID = deliveryPersonID
		order.EstimatedDelivery = estimatedDelivery
		order.Status = "En Route"
		order.DeliveryArea = deliveryArea

		if result := DB.Create(&order); result.Error != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON("Cannot Create Order")
		}
	}

	return c.RenderJSON(order)
}

// GetOrder returns a single order object
func (c *Orders) GetOrder(id int) revel.Result {
	order := models.Order{}
	result := DB.First(&order, id)
	println(result)
	return c.RenderJSON(order)
}
