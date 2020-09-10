package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

// Delivery Controller with embedded App
type Delivery struct {
	App
}

// GetDeliveryPeople is used to get all delivery people json
func (c *Delivery) GetDeliveryPeople() revel.Result {
	deliverPeople := []models.DeliveryPerson{}
	result := DB.Find(&deliverPeople)
	println(result)
	return c.RenderJSON(deliverPeople)
}

// GetDeliveryPerson is used to get a single delivery person object
func (c *Delivery) GetDeliveryPerson(id int) revel.Result {
	dp := models.DeliveryPerson{}
	result := DB.First(&dp, id)
	println(result.RowsAffected)
	return c.RenderJSON(dp)
}
