package models

import (
	"fmt"
	"strconv"
	"strings"
)

type KickAction struct {
	Name string
	Port int
	Sign string
}

func (c *KickAction) GetName() string {
	return c.Name
}

func (c *KickAction) GetParams(params map[string]string) {
	c.Name = params["a"]
	if port, err := strconv.Atoi(params["p"]); err == nil {
		c.Port = port
	}
}

func (c *KickAction) IsHope() bool {
	if c.Name == "kick" {
		return true
	}
	return false
}

func (c *KickAction) CheckParams() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("检查端口(必须在0-65535之间)")
	}
	return nil
}

func (c *KickAction) JoinPayload() *strings.Reader {
	return strings.NewReader(fmt.Sprintf("action=%s&password=%s&sign=%s", c.Name, Password, c.Sign))
}

func (c *KickAction) JoinUrl() string {
	return fmt.Sprintf(URL, c.Port)
}
