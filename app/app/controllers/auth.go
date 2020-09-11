package controllers

import (
	"app/app/models"
	"net/http"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"

	"app/app/services"
)

// Auth Controller with embedded App
type Auth struct {
	App
}

// Login creates a Token for user and Logs him in
func (c *Auth) Login() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	email := jsonData["email"].(string)
	password := jsonData["password"].(string)

	accessToken, refreshToken, err := services.LogIn(email, password)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Credentials"})
	}

	return c.RenderJSON(map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
}

// Register creates a new user and
func (c *Auth) Register() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	userType := jsonData["user_type"].(string)
	password := jsonData["password"].(string)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if userType == "customer" {
		obj := models.User{
			Profile: models.Profile{
				Email:    jsonData["email"].(string),
				Password: hashedPassword,
			},
		}
		result := services.DB.Create(&obj)
		println(result)
	} else {
		obj := models.DeliveryPerson{
			Profile: models.Profile{
				Email:    jsonData["email"].(string),
				Password: hashedPassword,
			},
		}
		result := services.DB.Create(&obj)
		println(result)
	}
	return c.RenderJSON(jsonData)
}

// RefreshToken returns accessToken using refreshToken
func (c *Auth) RefreshToken() revel.Result {
	if authToken := c.Request.Header.Get("Authorization"); authToken != "" {
		println("authToken", authToken)
		accessToken, err := services.GetRefreshToken(authToken)

		if err != nil {
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(map[string]string{"status": "Invalid Credentials"})
		}

		return c.RenderJSON(map[string]string{"accessToken": accessToken})
	}

	return c.RenderJSON(map[string]string{"status": "Invalid Credentials"})
}
