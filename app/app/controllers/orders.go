package controllers

import (
	"github.com/revel/revel"
)

type Orders struct {
	App
}

// Order Controllers
func (c *Orders) GetOrders() revel.Result {
	return c.RenderJSON(cookies)
}

func (c *Orders) CreateOrder() revel.Result {
	return c.RenderJSON(cookies)
}

func (c *Orders) GetOrder(id int) revel.Result {
	return c.RenderJSON(cookies)
}
