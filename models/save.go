package models

import (
	"fmt"
	"strings"
)

type SaveAction struct {
	Name     string
	Port     int
	ServerID int
}

func (c *SaveAction) IsHope() bool {
	if c.Name == "save" {
		return true
	}
	return false
}

func (c *SaveAction) GetName() string {
	return c.Name
}

func (c *SaveAction) GetParams(params map[string]interface{}) {
	if name, ok := params["a"].(string); ok {
		c.Name = name
	}
	if port, ok := params["p"].(int); ok {
		c.Port = port
	}
	if serverId, ok := params["s"].(int); ok {
		c.ServerID = serverId
	}
}

func (c *SaveAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	if c.ServerID == 0 {
		return fmt.Errorf("请输入ServerID")
	}
	return nil
}

func (c *SaveAction) JoinPayload() *strings.Reader {
	//payload := strings.NewReader(fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, SaveType, Md5([]byte(SecretKey))))
	return nil
}

func (c *SaveAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port, fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, SaveType, Md5(SecretKey)))
}
