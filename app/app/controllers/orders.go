package controllers

import (
	"net/http"

	"app/app/services"

	"github.com/revel/revel"
)

// Orders Controller with embedded App
type Orders struct {
	App
}

// GetOrders get all current orders
func (c *Orders) GetOrders() revel.Result {
	orders, err := services.QueryOrders()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// GetOrdersUser get all current orders of a user
func (c *Orders) GetOrdersUser(userid uint) revel.Result {
	orders, err := services.QueryOrdersUser(userid)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// GetOrdersDeliveryAgent get all current orders of delivery agent
func (c *Orders) GetOrdersDeliveryAgent(deliverypersonid uint) revel.Result {
	orders, err := services.QueryOrdersDeliveryAgent(deliverypersonid)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// CreateOrder create a new order
func (c *Orders) CreateOrder() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	order, err := services.InsertOrder(jsonData)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(order)
}

// GetOrder returns a single order object
func (c *Orders) GetOrder(id int) revel.Result {
	order, err := services.QueryOrder(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(order)
}
