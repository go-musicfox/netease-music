package service

import (
	"fmt"
	"testing"
)

func TestLoginCellphoneService_LoginCellphone(t *testing.T) {
	service := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	code, resp := service.LoginCellphone()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
