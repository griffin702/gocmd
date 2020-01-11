package models

import (
	"fmt"
	"strconv"
	"strings"
)

type HotAction struct {
	Name			string
	Port			int
	Version			string
	Sign			string
	Status			int
}

func (c *HotAction) GetName() string {
	return c.Name
}

func (c *HotAction) GetParams(params map[string]string) {
	c.Name = params["-a"]
	if port, err := strconv.Atoi(params["-p"]); err == nil {
		c.Port = port
	}
	c.Version = params["-v"]
}

func (c *HotAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	if c.Name == "hot" {
		return nil
	}
	return fmt.Errorf("非法参数")
}

func (c *HotAction) JoinPayload() *strings.Reader {
	return strings.NewReader(fmt.Sprintf("action=%s&password=%s&sign=%s&ver=%s", c.Name, Password, c.Sign, c.Version))
}

func (c *HotAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port)
}
