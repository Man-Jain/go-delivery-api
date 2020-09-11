package controllers

import (
	"net/http"

	"github.com/revel/revel"

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

// DeliveryLogin creates a Token for user and Logs him in
func (c *Auth) DeliveryLogin() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	email := jsonData["email"].(string)
	password := jsonData["password"].(string)

	accessToken, refreshToken, err := services.DeliveryLogIn(email, password)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Credentials"})
	}

	return c.RenderJSON(map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
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
