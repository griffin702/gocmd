package controllers

import "strings"

type Action interface {
	GetName()								string
	GetParams(params map[string]string)
	CheckParams()							error
	JoinPayload()							*strings.Reader
	JoinUrl()								string
}
