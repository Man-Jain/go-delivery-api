package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

// Users Controller with embedded App
type Users struct {
	App
}

// GetUsers return all users of application
func (c *Users) GetUsers() revel.Result {
	users := []models.User{}
	result := DB.Find(&users)
	println("the users are", result)
	println(result.RowsAffected)
	return c.RenderJSON(users)
}

// GetUser return a single user object
func (c *Users) GetUser(id int) revel.Result {
	user := models.User{}
	result := DB.First(&user, id)
	println(result.RowsAffected)
	return c.RenderJSON(user)
}
