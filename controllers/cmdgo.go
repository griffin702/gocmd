package controllers

import (
	"fmt"
	"gocmd/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type CmdGo struct {
	ParamList		map[string]string
	ActionList		*[]Action
	IsHelp			bool
}

func (c *CmdGo) ParseArgs(args []string) {
	for i := 1; i < len(args); i++ {
		key, value := args[i][:2], args[i][2:]
		switch key {
		case "-h":
			if len(key) == 2 && i == 1 {
				fmt.Println("帮助")
				c.IsHelp = true
			}
		default:
			c.ParamList[key] = value
		}
	}
}

func (c *CmdGo) RegistAction() {
	c.ActionList = &[]Action{
		new(models.KickAction),
		new(models.SaveAction),
		new(models.CloseAction),
		new(models.HotAction),
	}
}

func (c *CmdGo) SendRequest(url string, payload *strings.Reader) (num int, err error) {
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=gb2312")
	fmt.Println(req.URL, payload)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return
}

func (c *CmdGo) Run() error {
	c.RegistAction()
	for _, action := range *c.ActionList {
		action.GetParams(c.ParamList)
		if action.IsHope() {
			if err := action.CheckParams(); err != nil {
				return err
			}
			_, err := c.SendRequest(action.JoinUrl(), action.JoinPayload())
			if err != nil {
				return fmt.Errorf("发送[%s]请求>>Error：%s", action.GetName(), err.Error())
			}
		}
	}
	return nil
}