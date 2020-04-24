package commands

import (
	"gitee.com/griffin702/gocmd/models/flags"
	"github.com/urfave/cli/v2"
)

func InitCommands() []*cli.Command {
	f := new(flags.Flags)
	baseFlags := f.ToBaseFlags()
	baseAction := BaseAction(f)
	return []*cli.Command{
		{
			Name:     "close",
			Category: "BaseAction",
			Usage:    "关闭指定服务器",
			Flags:    baseFlags,
			Action:   baseAction,
		},
		{
			Name:     "kick",
			Category: "BaseAction",
			Usage:    "踢人操作",
			Flags:    baseFlags,
			Action:   baseAction,
		},
		{
			Name:     "hot",
			Category: "BaseAction",
			Usage:    "热更踢人操作",
			Flags:    baseFlags,
			Action:   baseAction,
		},
		{
			Name:     "save",
			Category: "BaseAction",
			Usage:    "保存数据操作",
			Flags:    baseFlags,
			Action:   baseAction,
		},
	}
}
