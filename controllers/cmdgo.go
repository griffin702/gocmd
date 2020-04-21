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
	help    bool
	action  string
	port    string
	version string
)

func init() {
	flag.BoolVar(&help, "h", false, "查看帮助")
	flag.StringVar(&action, "a", "save", "`action`：需要进行的操作")
	flag.StringVar(&port, "p", "10006", "`port`：指定请求端口")
	flag.StringVar(&version, "v", "", "`hot-version`：热更版本号")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `欢迎使用GOCMD
Usage: gocmd [-a action] [-p port] [-v hot-version]

Options:
`)
		flag.PrintDefaults()
	}
}

type CmdGo struct {
	ParamList  map[string]string
	ActionList *[]Action
	IsHelp     bool
}

type Action interface {
	GetName() string
	GetParams(params map[string]string)
	IsHope() bool
	CheckParams() error
	JoinPayload() *strings.Reader
	JoinUrl() string
}

func New() *CmdGo {
	cmdGo := new(CmdGo)
	cmdGo.IsHelp = help
	cmdGo.ParamList = make(map[string]string)
	cmdGo.ParamList["a"] = action
	cmdGo.ParamList["p"] = port
	cmdGo.ParamList["v"] = version
	cmdGo.RegistAction()
	return cmdGo
}

func (c *CmdGo) RegistAction() {
	c.ActionList = &[]Action{
		new(models.KickAction),
		new(models.SaveCloseAction),
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
