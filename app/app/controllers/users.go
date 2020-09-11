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
	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	users, err := services.QueryAllUsers()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(users)
}

// GetUser return a single user object
func (c *Users) GetUser(id uint) revel.Result {
	tid, _, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Auth Credentials"})
	}

	if id == tid {
		user, err := services.QueryUser(id)
		if err != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
		return c.RenderJSON(user)
	}

	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	user, err := services.QueryUser(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(user)
}

// Register creates a new user and
func (c *Users) Register() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	userType := jsonData["user_type"].(string)

	if userType == "customer" {
		_, err := services.InsertUser(jsonData)
		if err != nil {
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
	} else {
		_, err := services.InsertDeliveryPerson(jsonData)
		if err != nil {
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}
	}
	return c.RenderJSON(jsonData)
}

// RegisterRoot creates a new user and
func (c *Users) RegisterRoot() revel.Result {
	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	user, err := services.InsertRootUser(jsonData)
	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(user)
}
