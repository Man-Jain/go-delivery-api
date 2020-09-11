package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"app/app/services"
)

// Users Controller with embedded App
type Users struct {
	App
}

// GetUsers return all users of application
func (c *Users) GetUsers() revel.Result {
	users, err := services.QueryAllUsers()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(users)
}

// GetUser return a single user object
func (c *Users) GetUser(id int) revel.Result {
	user, err := services.QueryUser(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(user)
}
