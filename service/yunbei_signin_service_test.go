package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SigninResult struct {
	Code  int32  `json:"code"`
	Msg   string `json:"msg"`
	Point int32  `json:"point"`
}

func TestYunbeiSigninService_Signin(t *testing.T) {
	login := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	code, content, err := login.LoginCellphone()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	t.Logf("code: %f, content: %s", code, content)
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("content: %s", content)
	}
	service := &YunbeiSigninService{}
	code, resp, err := service.Signin()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	fmt.Println(code, string(resp))
	result := SigninResult{}
	_ = json.Unmarshal(resp, &result)
	if code != 200 || result.Point == 0{
		t.Errorf("签到失败！code: %f, msg: %s", code, result.Msg)
	}
}
