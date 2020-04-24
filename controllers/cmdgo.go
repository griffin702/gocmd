package controllers

import (
	"fmt"
	"gitee.com/griffin702/gocmd/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type CmdGo struct {
	Flags  *models.Flags
	Action models.Action
}

func New(flags *models.Flags) *CmdGo {
	cmdGo := new(CmdGo)
	cmdGo.Flags = flags
	cmdGo.ActionRegister()
	return cmdGo
}

func (c *CmdGo) ActionRegister() {
	c.Action = models.Actions{
		new(models.BaseAction),
	}
}

func (c *CmdGo) SendRequest(method, url string, payload *strings.Reader) (num int, err error) {
	m := strings.ToUpper(method)
	req, _ := http.NewRequest(m, url, nil)
	if m == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=gb2312")
	}
	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return
}

func (c *CmdGo) Run() (err error) {
	c.Action, err = c.Action.GetAction(c.Flags.Action)
	if err != nil {
		return err
	}
	c.Action.InitFlags(c.Flags)
	if err = c.Action.CheckParams(); err != nil {
		return
	}
	method, url := c.Action.JoinUrl()
	_, err = c.SendRequest(method, url, c.Action.JoinPayload())
	if err != nil {
		return fmt.Errorf("发送[%s]请求>>Error：%s", c.Action.GetName(), err.Error())
	}
	return nil
}
