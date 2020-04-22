package models

import (
	"gitee.com/griffin702/service/tools"
)

const (
	URL       = "http://127.0.0.1:%d/GoCMD?%s"
	SecretKey = "321321"
)

const (
	KickType = iota
	SaveType
	CloseType
	HotType
)

func Md5(str string) string {
	return tools.Tools.EncodeMD5(str)
}
