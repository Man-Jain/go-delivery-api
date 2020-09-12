package controllers

import (
	"github.com/revel/revel"
)

// App main controller
type App struct {
	*revel.Controller
}

// Index is for the root url
func (c *App) Index() revel.Result {
	return c.Render()
}
