package service

import (
	"testing"
)

func TestLoginCellphoneService_LoginCellphone(t *testing.T) {
	service := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	code, content, err := service.LoginCellphone()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	t.Logf("code: %f, content: %s", code, content)
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("content: %s", content)
	}
}
