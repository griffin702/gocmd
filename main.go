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
		if isHelp := app.Args(os.Args); isHelp {
			return
		}
		if err := app.Run(); err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}