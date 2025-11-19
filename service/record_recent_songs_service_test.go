package service

import (
	"testing"
)

func TestRecordRecentSongsService(t *testing.T) {
	login := &LoginCellphoneService{
		Phone:    "",
		Password: "",
	}
	_, _, err := login.LoginCellphone()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	recent_songs := &RecordRecentSongsService{}
	code, resp, err := recent_songs.RecordRecentSongs()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	t.Logf("code: %f, resp: %s", code, resp)
	if code != 200 {
		t.Errorf("code error: %f", code)
		t.Errorf("resp: %s", resp)
	}
}