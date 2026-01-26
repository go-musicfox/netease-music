package service

import (
	"testing"
)

func TestYunbeiSigninService_Signin(t *testing.T) {
	login := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	code, content, err := login.LoginCellphone()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	// t.Logf("code: %f, content: %s", code, content)
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("content: %s", content)
	}
	service := &YunbeiService{}
	sign, err := service.Sign()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	if sign.Code != 200 {
		t.Errorf("签到失败！code: %f", sign.Code)
	} else {
		if sign.Data.Sign == false && sign.Data.YunbeiNum == 0 {
			t.Logf("今天已经签到过啦!明天再来吧")
		} else {
			t.Logf("签到成功！获得云贝：%d 个", sign.Data.YunbeiNum)
		}
	}
	info, err := service.Info()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	if info.Code != 200 {
		t.Errorf("获取信息失败，code: %f", sign.Code)
	} else {
		t.Log("云贝信息：")
		t.Log("\n" + Format(info))
	}
}
