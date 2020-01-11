package models

import (
	"fmt"
	"strconv"
	"strings"
)

type CloseAction struct {
	Name			string
	Port			int
	Sign			string
	Status			int
}

func (c *CloseAction) GetName() string {
	return c.Name
}

func (c *CloseAction) GetParams(params map[string]string) {
	c.Name = params["-a"]
	if port, err := strconv.Atoi(params["-p"]); err == nil {
		c.Port = port
	}
}

func (c *CloseAction) IsHope() bool {
	if c.Name == "close" {
		return true
	}
	return false
}

func (c *CloseAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	return nil
}

func (c *CloseAction) JoinPayload() *strings.Reader {
	return strings.NewReader(fmt.Sprintf("action=%s&password=%s&sign=%s", c.Name, Password, c.Sign))
}

func (c *CloseAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port)
}
