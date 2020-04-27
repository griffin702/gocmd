package main

import (
	"github.com/griffin702/gocmd/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoCMD"
	app.Usage = "运维工具集"
	app.Version = getVersion()
	app.Authors = []*cli.Author{{
		Name:  "WuYun",
		Email: "117976509@qq.com",
	}}
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "查看帮助",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "GoCMD Version",
	}
	app.Commands = commands.InitCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}
