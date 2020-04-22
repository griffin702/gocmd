package models

import (
	"crypto/md5"
	"fmt"
)

const (
	URL       = "http://127.0.0.1:%d/GoCMD?%s"
	SecretKey = "321321"
)

const (
	KickType = iota
	SaveType
	CloseType
	HotType
)

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
