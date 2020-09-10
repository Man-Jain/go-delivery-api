package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

// Auth Controller with embedded App
type Auth struct {
	App
}

// Login creates a Token for user and Logs him in
func (c *Auth) Login() revel.Result {
	return c.RenderJSON("cookies")
}

// Register creates a new user and
func (c *Auth) Register() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	userType := jsonData["user_type"].(string)
	if userType == "customer" {
		obj := models.User{
			Profile: models.Profile{
				Email:    jsonData["email"].(string),
				Password: jsonData["password"].(string),
			},
		}
		result := DB.Create(&obj)
		println(result)
	} else {
		obj := models.DeliveryPerson{
			Profile: models.Profile{
				Email:    jsonData["email"].(string),
				Password: jsonData["password"].(string),
			},
		}
		result := DB.Create(&obj)
		println(result)
	}
	return c.RenderJSON(jsonData)
}
