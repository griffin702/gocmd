package commands

import (
	"github.com/griffin702/gocmd/controllers"
	"github.com/griffin702/gocmd/models/flags"
	"github.com/urfave/cli/v2"
)

func BaseAction(f *flags.Flags) cli.ActionFunc {
	return func(ctx *cli.Context) (err error) {
		f.Name = ctx.Command.Name
		cmdGo := controllers.New(f)
		return cmdGo.Run()
	}
}
