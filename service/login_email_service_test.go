package service

import (
	"fmt"
	"testing"
)

func TestLoginEmailService_LoginCellphone(t *testing.T) {
	service := &LoginEmailService{
		Email:    "",
		Password: "",
	}
	code, resp := service.LoginEmail()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
