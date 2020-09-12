package controllers

import (
	"net/http"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"

	"app/app/models"
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

// Register creates a new user
func (c *Users) Register() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	userType, oku := jsonData["user_type"].(string)
	password, okp := jsonData["password"].(string)
	email, oke := jsonData["email"].(string)
	if !oku || !okp || !oke {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if userType == "customer" {
		obj := models.User{
			Profile: models.Profile{
				Email:    email,
				Password: hashedPassword,
			},
		}

		obj.Validate(c.Validation)
		if c.Validation.HasErrors() {
			return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
		}

		_, err := services.InsertUser(obj)
		if err != nil {
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}

		return c.RenderJSON(obj)

	} else {

		obj := models.DeliveryPerson{
			Profile: models.Profile{
				Email:    email,
				Password: hashedPassword,
			},
		}
		obj.Validate(c.Validation)
		if c.Validation.HasErrors() {
			return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
		}
		_, err := services.InsertDeliveryPerson(obj)
		if err != nil {
			return c.RenderJSON(map[string]string{"status": "Invalid Request"})
		}

		return c.RenderJSON(obj)
	}
}

// RegisterRoot creates a new root user
func (c *Users) RegisterRoot() revel.Result {
	isAllowed, err := services.ValidateUser(c.Request.Header.Get("Authorization"), "root")
	if !isAllowed {
		return c.RenderJSON(map[string]string{"status": "Not Authorised"})
	}

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	email, okEmail := jsonData["email"].(string)
	password, okPass := jsonData["password"].(string)
	if !okEmail || !okPass {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	obj := models.User{
		Profile: models.Profile{
			Email:    email,
			Password: hashedPassword,
		},
		Root: true,
	}

	_, err = services.InsertRootUser(obj)
	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	if err != nil {
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(obj)
}
