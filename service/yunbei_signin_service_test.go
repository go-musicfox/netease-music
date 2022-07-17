package service

import (
	"fmt"
	"testing"
)

func TestYunbeiSigninService_Signin(t *testing.T) {
	service := &YunbeiSigninService{}
	code, resp := service.Signin()
	fmt.Println(code, string(resp))
	if code != 301 {
		t.Errorf("code error: %f", code)
	}
}
