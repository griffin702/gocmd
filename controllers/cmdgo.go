package controllers

import (
	"fmt"
	"strconv"
)

type CmdGo struct {
	ParamList		map[string]string
}

func (c *CmdGo) Args(args []string) bool {
	for i := 1; i < len(args); i++ {
		key, value := args[i][:2], args[i][2:]
		switch key {
		case "-h":
			if len(key) == 2 && i == 1 {
				fmt.Println("帮助")
				return true
			}
		default:
			c.ParamList[key] = value
		}
	}
	return false
}

func (c *CmdGo) IsInvalid() (bool, error) {
	action := c.ParamList["-a"]
	port, err := strconv.Atoi(c.ParamList["-p"])
	if err != nil || port < 0 || port > 65535 {
		return true, fmt.Errorf("必须输入端口(0-65535)")
	}
	if action == "hot" && c.ParamList["-v"] == "" {
		return true, fmt.Errorf("参数不合法")
	}
	if action != "" {
		return false, nil
	}
	return true, fmt.Errorf("参数不合法")
}

func (c *CmdGo) Run() error{
	if ok, err := c.IsInvalid(); ok { return err }
	action := c.ParamList["-a"]
	ver := c.ParamList["-v"]
	port := c.ParamList["-p"]
	worker := &Worker{}
	worker.Url = fmt.Sprintf("http://127.0.0.1:%s/gocmd", port)
	worker.Password = "321321"
	worker.Action = action
	worker.Ver = ver
	switch action {
	case "kick":
		_, err := worker.SendRequest()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("踢人动作返回消息")
	case "save":
		_, err := worker.SendRequest()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("保存动作返回消息")
	case "close":
		_, err := worker.SendRequest()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("关服动作返回消息")
	case "hot":
		_, err := worker.SendRequest()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("热更动作返回消息")
	default:
		return fmt.Errorf("参数不合法")
	}
	return nil
}