package main

import (
	"flag"
	"fmt"
	"gocmd/controllers"
)

func main() {
	flag.Parse()
	app := controllers.New()
	if app.IsHelp {
		flag.Usage()
		return
	}
	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
