package models

import (
	"fmt"
	"strconv"
	"strings"
)

type SaveAction struct {
	Name			string
	Port			int
	Sign			string
}

func (c *SaveAction) GetName() string {
	return c.Name
}

func (c *SaveAction) GetParams(params map[string]string) {
	c.Name = params["-a"]
	if port, err := strconv.Atoi(params["-p"]); err == nil {
		c.Port = port
	}
}

func (c *SaveAction) IsHope() bool {
	if c.Name == "save" {
		return true
	}
	return false
}

func (c *SaveAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	return nil
}

func (c *SaveAction) JoinPayload() *strings.Reader {
	return strings.NewReader(fmt.Sprintf("action=%s&password=%s&sign=%s", c.Name, Password, c.Sign))
}

func (c *SaveAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port)
}
