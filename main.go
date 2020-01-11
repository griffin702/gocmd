package main

import (
	"fmt"
	"gocmd/controllers"
	"os"
)

func main() {
	// name := config.Conf.Get("app.name").(string)
	app := controllers.App
	argNum := len(os.Args)
	if argNum > 1 {
		app.ParseArgs(os.Args)
		if app.IsHelp { return }
		if err := app.Run(); err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}