package action

import (
	"fmt"
	"gitee.com/griffin702/gocmd/models/flags"
	"gitee.com/griffin702/service/tools"
	"strings"
)

const (
	SecretKey = "321321"
	BaseURL   = "http://%s:%d/GoCMD?%s"
)

const (
	BaseActionType = iota
)

const (
	_ = iota
	CloseType
	KickType
	HotType
	SaveType
)

var (
	actionTypeMap = map[string][]int{
		"close": {BaseActionType, CloseType},
		"kick":  {BaseActionType, KickType},
		"hot":   {BaseActionType, HotType},
		"save":  {BaseActionType, SaveType},
	}
)

type Action interface {
	GetName() string
	GetType() int
	GetAction(action string) (Action, error)
	InitFlags(flags *flags.Flags)
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

func (c Actions) InitFlags(flags *flags.Flags) {
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

func Md5(str string) string {
	return tools.Tools.EncodeMD5(str)
}
