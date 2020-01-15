package models

import "gitee.com/griffin702/services/tool"

var tools = tool.Tools

const (
	URL = "http://127.0.0.1:%d/gocmd"
	Password = "321321"
	//SecretKey = ""
)

const (
	helpStr = ` -a		期望执行的action，当前支持'kick','save','close','hot'
 -p		目标接口所属的端口，必填
 -h		帮助
 -ahot -v	热更action需填入版本号`
	kickStr = ` -akick		执行踢人action
 -p		目标接口所属的端口，必填`
	saveCloseStr = ` -asaveclose		执行安全关服action
 -p		目标接口所属的端口，必填`
	hotStr = ` -ahot		执行热更action，-v必填参数
 -v		目标版本号
 -p		目标接口所属的端口，必填`
)