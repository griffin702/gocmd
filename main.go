package main

import (
	"fmt"
	"gitee.com/griffin702/gocmd/controllers"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	af controllers.ActionFlags
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
			Usage:       "指定操作`行为`",
			Destination: &af.Action,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			Value:       "127.0.0.1",
			Usage:       "指定服务器`IP`地址",
			Destination: &af.Host,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "指定服务器`端口`",
			Destination: &af.Port,
			Required:    true,
		},
		&cli.IntFlag{
			Name:        "server",
			Aliases:     []string{"s"},
			Usage:       "指定服务器`ID`",
			Destination: &af.ServerID,
			Required:    true,
		},
	}
	app.Action = func(ctx *cli.Context) (err error) {
		cmdGo := controllers.New(&af)
		err = cmdGo.Run(af.Action)
		return
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
