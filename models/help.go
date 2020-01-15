package models

import "fmt"

type Help struct {
	Content			*map[string]string
}

func (c *Help) New() {
	c.Content = &map[string]string{
		"help": helpStr,
		"kick": kickStr,
		"save": saveStr,
		"close": closeStr,
		"hot": hotStr,
	}
}

func (c *Help) ShowContent(name string) {
	content := *c.Content
	fmt.Println(content[name])
	if content[name] == "" {
		fmt.Println("帮助使用不正确")
	}
}
