package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"app/app/models"
	"app/app/services"
)

// Cookies Controller with embedded App
type Cookies struct {
	App
}

// GetCookies is used to get all cookies json
func (c *Cookies) GetCookies() revel.Result {
	cookies, err := services.QueryAllCookies()
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(cookies)
}

// CreateCookie is used to create a new cookie item
func (c *Cookies) CreateCookie() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	name, okname := jsonData["name"].(string)
	description, okdesc := jsonData["description"].(string)
	price, okprice := jsonData["price"].(float64)
	quantity, okquantity := jsonData["quantity"].(float64)

	if !okname || !okdesc || !okprice || !okquantity {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	cookie := models.Cookie{
		Name:        name,
		Description: description,
		Price:       uint(price),
		Quantity:    uint(quantity),
	}

	cookie.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderJSON(map[string]string{"status": "Invalid Parameters"})
	}

	_, err := services.InsertCookie(cookie)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}

	return c.RenderJSON(jsonData)
}

// GetCookie is used to get a single cookie object
func (c *Cookies) GetCookie(id int) revel.Result {
	cookie, err := services.QueryCookie(id)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{"status": "Invalid Request"})
	}
	return c.RenderJSON(cookie)
}
