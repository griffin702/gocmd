package controllers

import (
	"gitee.com/griffin702/services"
	"gocmd/models"
)

var (
	App = newApp()
)

func newApp() (app *Gocmd.CmdGo) {
	app = Gocmd.New()
	app.RegistAction([]Gocmd.Action{
		new(models.KickAction),
		new(models.SaveAction),
		new(models.CloseAction),
		new(models.HotAction),
	})
	return
}