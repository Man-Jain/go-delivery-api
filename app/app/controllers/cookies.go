package controllers

import (
	"github.com/revel/revel"
)

type Cookies struct {
	App
}

var cookies = []Cookie{
	Cookie{"Chocolate", "This is chocolate cookie", 123},
	Cookie{"Cranberry", "This is Cranberry cookie", 123},
	Cookie{"Blueberry", "This is Blueberry cookie", 12123},
}

// Cookies Controllers
func (c *Cookies) GetCookies() revel.Result {
	return c.RenderJSON(cookies)
}

func (c *Cookies) CreateCookie() revel.Result {
	return c.RenderJSON(cookies)
}

func (c *Cookies) GetCookie(id int) revel.Result {
	return c.RenderJSON(cookies)
}
