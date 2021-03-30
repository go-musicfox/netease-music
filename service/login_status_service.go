package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"github.com/anhoder/netease-music/util"
	"strings"
)

type LoginStatusService struct{}

func (service *LoginStatusService) LoginStatus() map[string]interface{} {

	options := &util.Options{
		Crypto:  "",
		Ua:      "",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("GET", `https://music.163.com`, data, options)

	bodyText := reBody["html"].(string)
	bodyJson1 := make(map[string]interface{})
	var bodyJson2 []interface{}
	reg1 := regexp.MustCompile(`GUser\s*=\s*([^;]+);`)
	reg2 := regexp.MustCompile(`GBinds\s*=\s*([^;]+);`)
	s1 := []rune(reg1.FindStringSubmatch(bodyText)[0])

	// 如果此字符过短，说明用户没有登录
	if len(s1) < 12 {
		delete(reBody, "html")
		reBody["code"] = 301
		reBody["msg"] = "需要登录"
		return reBody
	}

	s2 := []rune(reg2.FindStringSubmatch(bodyText)[0])
	result1 := string(s1[6 : len(s1)-1])
	result1 = strings.Replace(result1, "{", "{\"", -1)
	result1 = strings.Replace(result1, "http://", "http", -1)
	result1 = strings.Replace(result1, ":", "\":", -1)
	result1 = strings.Replace(result1, ",", ",\"", -1)
	result1 = strings.Replace(result1, "http", "http://", -1)
	result2 := string(s2[7 : len(s2)-1])
	fmt.Println(result1)
	fmt.Println(result2)
	err1 := json.Unmarshal([]byte(result1), &bodyJson1)
	err2 := json.Unmarshal([]byte(result2), &bodyJson2)
	delete(reBody, "html")
	if err1 != nil || err2 != nil {
		reBody["code"] = 500
		if err1 == nil {
			err1 = errors.New("nil")
		}
		if err2 == nil {
			err2 = errors.New("nil")
		}
		reBody["err"] = err1.Error() + "\n=========\n" + err2.Error()
	} else {
		reBody["profile"] = bodyJson1
		reBody["bindings"] = bodyJson2
	}
	return reBody
}
