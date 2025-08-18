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
	t.Logf("code: %f, content: %s", code, content)
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("content: %s", content)
	}
}
