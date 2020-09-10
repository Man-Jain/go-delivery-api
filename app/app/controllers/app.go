package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Cookie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c *App) Index() revel.Result {
	return c.Render()
}
