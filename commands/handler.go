package commands

import (
	"github.com/griffin702/gocmd/controllers"
	"github.com/griffin702/gocmd/models/flags"
	"github.com/urfave/cli/v2"
)

// 使用request请求服务端接口
func BaseAction(f *flags.Flags) cli.ActionFunc {
	return func(ctx *cli.Context) (err error) {
		f.Name = ctx.Command.Name
		cmdGo := controllers.NewCmdGo(f)
		return cmdGo.Run()
	}
}
