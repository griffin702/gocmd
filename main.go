package main

import (
	"fmt"
	"gitee.com/griffin702/gocmd/controllers"
	"gitee.com/griffin702/gocmd/models"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	flags models.Flags
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
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "action",
			Aliases:     []string{"a"},
			Value:       "save",
			Usage:       "指定操作`行为`",
			Destination: &flags.Action,
		},
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			Value:       "127.0.0.1",
			Usage:       "指定服务器`IP`地址",
			Destination: &flags.Host,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "指定服务器`端口`",
			Destination: &flags.Port,
		},
		&cli.IntFlag{
			Name:        "server",
			Aliases:     []string{"s"},
			Usage:       "指定服务器`ID`",
			Destination: &flags.ServerID,
		},
	}
	app.Action = func(ctx *cli.Context) (err error) {
		cmdGo := controllers.New(&flags)
		return cmdGo.Run()
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
