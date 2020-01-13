package controllers

import (
	"gitee.com/griffin702/services"
	"gocmd/models"
)

var (
	App = newApp()
)

func newApp() (app *services.CmdGo) {
	app = services.New()
	app.RegistAction([]services.Action{
		new(models.KickAction),
		new(models.SaveAction),
		new(models.CloseAction),
		new(models.HotAction),
	})
	return
}