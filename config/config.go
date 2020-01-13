package config

import (
	"fmt"
	"os"
	"strings"
	"github.com/pelletier/go-toml"
)

var (
	Conf = New()
)

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	conf := "./config/config.toml"
	if getAppEnv() == "-test" {
		conf = "./../config/config.toml"
	}
	config, err := toml.LoadFile(conf)
	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}

func getAppEnv() string {
	file := os.Args
	var s string
	if len(file) > 1 {
		s = strings.Split(file[1], ".")[0]
	}
	return s
}
