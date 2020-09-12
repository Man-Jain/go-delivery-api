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
	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	orders, err := services.QueryOrders()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// GetOrdersUser get all current orders of a user
func (c *Orders) GetOrdersUser(userid uint) revel.Result {
	tid, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	if userid == tid {
		orders, err := services.QueryOrdersUser(userid)
		if err != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
		return c.RenderJSON(orders)
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	orders, err := services.QueryOrdersUser(userid)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// GetOrdersDeliveryAgent get all current orders of delivery agent
func (c *Orders) GetOrdersDeliveryAgent(deliverypersonid uint) revel.Result {
	tid, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	if deliverypersonid == tid {
		orders, err := services.QueryOrdersDeliveryAgent(deliverypersonid)
		if err != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
		return c.RenderJSON(orders)
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	orders, err := services.QueryOrdersDeliveryAgent(deliverypersonid)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(orders)
}

// CreateOrder create a new order
func (c *Orders) CreateOrder() revel.Result {
	_, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	userID, okuid := jsonData["user_id"].(float64)
	cookieID, okcid := jsonData["cookie_id"].(float64)
	quantity, okqty := jsonData["quantity"].(float64)

	if !okuid || !okcid || !okqty {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	order, err := services.InsertOrder(uint(userID), uint(cookieID), uint(quantity))
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(order)
}

// GetOrder returns a single order object
func (c *Orders) GetOrder(id int) revel.Result {
	tid, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	order, err := services.QueryOrder(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	if order.UserID == tid {
		return c.RenderJSON(order)
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	return c.RenderJSON(order)
}
