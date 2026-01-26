package service

import (
	"testing"
)

func TestUserRecordService(t *testing.T) {
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
	record_service := &UserRecordService{
		UId: "",
	}
	code, resp := record_service.UserRecord()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	t.Logf("code: %f, content: %s", code, string(resp))
}
