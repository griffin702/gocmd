package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Worker struct {
	Url			string
	Action		string
	Ver			string
	Password	string
	Sign		string
}

func (c *Worker) SendRequest() (num int, err error) {
	payload := strings.NewReader(fmt.Sprintf("action=%s&password=%s&sign=%s&ver=%s", c.Action, c.Password, c.Sign, c.Ver))
	req, _ := http.NewRequest("POST", c.Url, payload)
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