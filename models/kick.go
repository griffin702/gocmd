package models

import (
	"fmt"
	"strings"
)

type KickAction struct {
	Name     string
	Port     int
	ServerID int
}

func (c *KickAction) IsHope() bool {
	if c.Name == "kick" {
		return true
	}
	return false
}

func (c *KickAction) GetName() string {
	return c.Name
}

func (c *KickAction) GetParams(params map[string]interface{}) {
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

func (c *KickAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	if c.ServerID == 0 {
		return fmt.Errorf("请输入ServerID")
	}
	return nil
}

func (c *KickAction) JoinPayload() *strings.Reader {
	//payload := strings.NewReader(fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, SaveType, Md5([]byte(SecretKey))))
	return nil
}

func (c *KickAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port, fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, KickType, Md5(SecretKey)))
}
