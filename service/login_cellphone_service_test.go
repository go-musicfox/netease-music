package service

import (
	"testing"
)

func TestLoginCellphoneService_LoginCellphone(t *testing.T) {
	service := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	code, content := service.LoginCellphone()
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("content: %s", content)
	}
}
