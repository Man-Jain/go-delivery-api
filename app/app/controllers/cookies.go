package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

// Cookies Controller with embedded App
type Cookies struct {
	App
}

var cookies = []Cookie{
	Cookie{"Chocolate", "This is chocolate cookie", 123},
	Cookie{"Cranberry", "This is Cranberry cookie", 123},
	Cookie{"Blueberry", "This is Blueberry cookie", 12123},
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

	cookie := models.Cookie{
		Name:        jsonData["name"].(string),
		Description: jsonData["description"].(string),
		Price:       jsonData["price"].(uint),
		Quantity:    jsonData["quantity"].(uint),
	}

	result := DB.Create(&cookie)
	println(result)

	return c.RenderJSON(cookie)
}

// GetCookie is used to get a single cookie object
func (c *Cookies) GetCookie(id int) revel.Result {
	cookie := models.Cookie{}
	result := DB.First(&cookie, id)
	println(result.RowsAffected)
	return c.RenderJSON(cookie)
}
