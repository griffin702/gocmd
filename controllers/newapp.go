package controllers

var (
	App = newApp()
)

func newApp() (app *CmdGo) {
	app = New()
	app.RegistAction()
	return
}
