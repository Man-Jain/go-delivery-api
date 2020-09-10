package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

// Cookies Controller with embedded App
type Cookies struct {
	App
}

// GetCookies is used to get all cookies json
func (c *Cookies) GetCookies() revel.Result {
	cookies := []models.Cookie{}
	result := DB.Find(&cookies)
	println(result)
	return c.RenderJSON(cookies)
}

// CreateCookie is used to create a new cookie item
func (c *Cookies) CreateCookie() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	println("This is the data", jsonData["name"])
	cookie := models.Cookie{
		Name:        jsonData["name"].(string),
		Description: jsonData["description"].(string),
		Price:       uint(jsonData["price"].(float64)),
		Quantity:    uint(jsonData["quantity"].(float64)),
	}

	result := DB.Create(&cookie)
	println(result.RowsAffected)

	return c.RenderJSON(jsonData)
}

// GetCookie is used to get a single cookie object
func (c *Cookies) GetCookie(id int) revel.Result {
	cookie := models.Cookie{}
	result := DB.First(&cookie, id)
	println(result.RowsAffected)
	return c.RenderJSON(cookie)
}
