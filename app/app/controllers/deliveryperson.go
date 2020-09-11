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
	deliveryPeople, err := services.QueryAllDeliveryPeople()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(deliveryPeople)
}

// GetDeliveryPerson is used to get a single delivery person object
func (c *Delivery) GetDeliveryPerson(id int) revel.Result {
	dp, err := services.QueryDeliveryPerson(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(dp)
}

// CompleteDelivery is used to get a single delivery person object
func (c *Delivery) CompleteDelivery(id int) revel.Result {
	_, err := services.UpdateDelivery(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON("Successfully Delivered")
}
