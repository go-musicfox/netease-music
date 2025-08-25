package service_test

import (
	"testing"

	"github.com/go-musicfox/netease-music/service"
)

func TestCaptchaSentService(t *testing.T) {
	service := &service.CaptchaSentService{
		Cellphone: "",
	}
	code, bodyBytes := service.CaptchaSent()
	t.Logf("code: %f, body: %s", code, string(bodyBytes))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
	if len(bodyBytes) == 0 {
		t.Error("response body is empty")
	}
}
