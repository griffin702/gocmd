package models

import (
	"fmt"
	"gitee.com/griffin702/service/tools"
	"strings"
)

const (
	SecretKey = "321321"
	BaseURL   = "http://127.0.0.1:%d/GoCMD?%s"
)

const (
	BaseActionType = iota
)

const (
	CloseType = 1
	KickType  = 2
	HotType   = 3
	SaveType  = 4
)

var (
	actionTypeMap = map[string][]int{
		"kick":  {BaseActionType, KickType},
		"save":  {BaseActionType, SaveType},
		"close": {BaseActionType, CloseType},
		"hot":   {BaseActionType, HotType},
	}
)

func Md5(str string) string {
	return tools.Tools.EncodeMD5(str)
}

type Action interface {
	GetName() string
	GetType() int
	GetAction(action string) (Action, error)
	GetParams(params map[string]interface{})
	CheckParams() error
	JoinPayload() *strings.Reader
	JoinUrl() (string, string)
}

type Actions []Action

func (c Actions) GetName() string {
	return ""
}

func (c Actions) GetType() int {
	return -1
}

func (c Actions) GetAction(action string) (Action, error) {
	t, ok := actionTypeMap[action]
	if !ok {
		return nil, fmt.Errorf("未找到action：%s", action)
	}
	for _, a := range c {
		if a.GetType() == t[0] {
			return a, nil
		}
	}
	return nil, fmt.Errorf("actionType error：%d", t)
}

func (c Actions) GetParams(params map[string]interface{}) {
	return
}

func (c Actions) CheckParams() error {
	return nil
}

func (c Actions) JoinPayload() *strings.Reader {
	return nil
}

func (c Actions) JoinUrl() (string, string) {
	return "", ""
}
