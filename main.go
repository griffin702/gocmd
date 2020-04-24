package main

import (
	"gitee.com/griffin702/gocmd/commands"
	"gitee.com/griffin702/gocmd/models/flags"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoCMD"
	app.Usage = "GoCMD运维工具集"
	app.Version = getVersion()
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "查看帮助",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "GoCMD Version",
	}
	var f flags.Flags
	app.Commands = commands.InitCommands(&f)
	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}
