package controllers

import (
	"github.com/revel/revel"
)

type Cookie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type App struct {
	*revel.Controller
}

var cookies = []Cookie{
	Cookie{"Chocolate", "This is chocolate cookie", 123},
	Cookie{"Cranberry", "This is Cranberry cookie", 123},
	Cookie{"Blueberry", "This is Blueberry cookie", 12123},
}

func (c App) Index() revel.Result {
	return c.Render()
}

// Cookies Controllers
func (c App) GetCookies() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) CreateCookie() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) GetCookie(id int) revel.Result {
	return c.RenderJSON(cookies)
}

// User Controllers
func (c App) Login() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) Register() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) GetUser(id int) revel.Result {
	return c.RenderJSON(cookies)
}

// Order Controllers
func (c App) GetOrders() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) CreateOrder() revel.Result {
	return c.RenderJSON(cookies)
}

func (c App) GetOrder(id int) revel.Result {
	return c.RenderJSON(cookies)
}
