package service

import (
	"fmt"
	"testing"
)

func TestLoginQRService(t *testing.T) {
	service := &LoginQRService{}
	code, resp, url := service.GetKey()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Fatalf("code error: %f", code)
	}
	if url == "" {
		t.Fatalf("url empty: %s", url)
	}

	code, resp = service.CheckQR()
	if code != 803 {
		t.Fatalf("code error: %f", code)
	}

	accountService := &UserAccountService{}
	code, resp = accountService.AccountInfo()
	if code != 200 {
		t.Fatalf("code error: %f", code)
	}
}
