package controllers

import (
	"flag"
	"fmt"
	"gocmd/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	help     bool
	action   string
	port     int
	serverId int
)

func init() {
	flag.BoolVar(&help, "h", false, "查看帮助")
	flag.StringVar(&action, "a", "save", "`action`：需要进行的操作")
	flag.IntVar(&port, "p", 0, "`port`：指定请求端口")
	flag.IntVar(&serverId, "s", 0, "`serverId`：服务端ID")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `欢迎使用GOCMD
Usage: gocmd [-a action] [-p port] [-s serverId] [-h help]

Options:
`)
		flag.PrintDefaults()
	}
}

type CmdGo struct {
	ParamList map[string]interface{}
	Action    models.Action
	IsHelp    bool
}

func New() *CmdGo {
	cmdGo := new(CmdGo)
	cmdGo.IsHelp = help
	cmdGo.ParamList = make(map[string]interface{})
	cmdGo.ParamList["a"] = action
	cmdGo.ParamList["p"] = port
	cmdGo.ParamList["s"] = serverId
	cmdGo.RegistAction()
	return cmdGo
}

func (c *CmdGo) RegistAction() {
	c.Action = models.Actions{
		new(models.BaseAction),
	}
}

func (c *CmdGo) SendRequest(url string, payload *strings.Reader) (num int, err error) {
	req, _ := http.NewRequest("GET", url, nil)
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

func (c *CmdGo) Run() (err error) {
	c.Action, err = c.Action.GetAction(action)
	if err != nil {
		return err
	}
	c.Action.GetParams(c.ParamList)
	if err = c.Action.CheckParams(); err != nil {
		return
	}
	_, err = c.SendRequest(c.Action.JoinUrl(), c.Action.JoinPayload())
	if err != nil {
		return fmt.Errorf("发送[%s]请求>>Error：%s", c.Action.GetName(), err.Error())
	}
	return nil
}
