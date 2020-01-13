package controllers

var (
	App = newApp()
)

func newApp() (app *CmdGo) {
	app = &CmdGo{
		ParamList: map[string]string{
			"-a": "",
			"-v": "",
		},
	}
	return
}