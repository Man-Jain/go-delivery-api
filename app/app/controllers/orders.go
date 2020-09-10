package controllers

import (
	"app/app/models"

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
	result := DB.Where("userid = ?", userid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return c.RenderJSON(orders)
}

// GetOrdersDeliveryAgent get all current orders of delivery agent
func (c *Orders) GetOrdersDeliveryAgent(deliverypersonid uint) revel.Result {
	orders := []models.Order{}
	result := DB.Where("deliverypersonid = ?", deliverypersonid).Find(&orders)
	println("the orders are", result)
	println(result.RowsAffected)
	return c.RenderJSON(orders)
}

// CreateOrder create a new order
func (c *Orders) CreateOrder() revel.Result {
	return c.RenderJSON(cookies)
}

// GetOrder returns a single order object
func (c *Orders) GetOrder(id int) revel.Result {
	order := models.Order{}
	result := DB.First(&order, id)
	println(result)
	return c.RenderJSON(order)
}
