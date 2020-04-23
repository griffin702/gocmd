package models

import (
	"fmt"
	"strings"
)

type BaseAction struct {
	Name     string
	Host     string
	Port     int
	ServerID int
}

func (c *BaseAction) GetType() int {
	return BaseActionType
}

func (c *BaseAction) GetName() string {
	return c.Name
}

func (c *BaseAction) GetAction(action string) (Action, error) {
	return c, nil
}

func (c *BaseAction) GetParams(params map[string]interface{}) {
	if name, ok := params["a"].(string); ok {
		c.Name = name
	}
	if host, ok := params["h"].(string); ok {
		c.Host = host
	}
	if port, ok := params["p"].(int); ok {
		c.Port = port
	}
	if serverId, ok := params["s"].(int); ok {
		c.ServerID = serverId
	}
}

func (c *BaseAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	if c.ServerID == 0 {
		return fmt.Errorf("请输入ServerID")
	}
	return nil
}

func (c *BaseAction) JoinPayload() *strings.Reader {
	//payload := strings.NewReader(fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, SaveType, Md5([]byte(SecretKey))))
	return nil
}

func (c *BaseAction) JoinUrl() (string, string) {
	return "get", fmt.Sprintf(BaseURL, c.Host, c.Port, fmt.Sprintf("ServerID=%d&Opt=%d&Sign=%s", c.ServerID, actionTypeMap[c.Name][1], Md5(SecretKey)))
}
