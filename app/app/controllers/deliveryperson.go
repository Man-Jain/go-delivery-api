package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"app/app/services"
)

// Delivery Controller with embedded App
type Delivery struct {
	App
}

// GetDeliveryPeople is used to get all delivery people json
func (c *Delivery) GetDeliveryPeople() revel.Result {
	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	deliveryPeople, err := services.QueryAllDeliveryPeople()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(deliveryPeople)
}

// GetDeliveryPerson is used to get a single delivery person object
func (c *Delivery) GetDeliveryPerson(id uint) revel.Result {
	tid, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	if id == tid {
		dp, err := services.QueryDeliveryPerson(id)
		if err != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
		return c.RenderJSON(dp)
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	dp, err := services.QueryDeliveryPerson(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(dp)
}

// CompleteDelivery is used to get a single delivery person object
func (c *Delivery) CompleteDelivery(id int) revel.Result {
	_, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "delivery_person")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	_, err = services.UpdateDelivery(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON("Successfully Delivered")
}
