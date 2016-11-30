package controllers

import (
	"github.com/revel/revel"
	"github.com/ledoerr/togo/gocockpit/app"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	services := app.GetAllServices()

	return c.Render(services)
}
