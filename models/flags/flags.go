package flags

import (
	"github.com/urfave/cli/v2"
)

type Flags struct {
	Name     string
	Host     string
	Port     int
	ServerID int
}

func (f *Flags) ToBaseFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			Value:       "127.0.0.1",
			Usage:       "指定服务器`IP`",
			Destination: &f.Host,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "指定服务`端口`",
			Destination: &f.Port,
			Required:    true,
		},
		&cli.IntFlag{
			Name:        "server",
			Aliases:     []string{"s"},
			Usage:       "指定`ServerID`",
			Destination: &f.ServerID,
			Required:    true,
		},
	}
}
