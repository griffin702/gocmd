package models

import (
	"fmt"
	"strings"
)

type CloseAction struct {
	Name     string
	Port     int
	ServerID int
}

func (c *CloseAction) IsHope() bool {
	if c.Name == "close" {
		return true
	}
	return false
}

func (c *CloseAction) GetName() string {
	return c.Name
}

func (c *CloseAction) GetParams(params map[string]interface{}) {
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

func (c *CloseAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	if c.ServerID == 0 {
		return fmt.Errorf("请输入ServerID")
	}
	return nil
}

func (c *CloseAction) JoinPayload() *strings.Reader {
	//payload := strings.NewReader(fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, SaveType, Md5([]byte(SecretKey))))
	return nil
}

func (c *CloseAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port, fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, CloseType, Md5(SecretKey)))
}
