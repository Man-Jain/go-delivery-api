package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

type Users struct {
	App
}

// User Controllers
func (c *Users) Login() revel.Result {
	return c.RenderJSON(cookies)
}

func (c *Users) Register() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	user := models.User{
		Profile: models.Profile{
			Email:    jsonData["email"].(string),
			Password: jsonData["password"].(string),
		},
	}

	result := DB.Create(&user)
	println(result)
	return c.RenderJSON(jsonData)
}

func (c *Users) GetUsers() revel.Result {
	users := []models.User{}
	result := DB.Find(&users)
	println("the users are", result)
	println(result.RowsAffected)
	return c.RenderJSON(users)
}

func (c *Users) GetUser(id int) revel.Result {
	user := models.User{}
	result := DB.First(&user, id)
	println(result.RowsAffected)
	return c.RenderJSON(user)
}
