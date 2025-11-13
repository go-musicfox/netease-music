package service

import (
	"fmt"
	"testing"
)

func TestSongUrlV1Service_SongUrl(t *testing.T) {
	service := &SongUrlV1Service{
		ID:    "405998841",
		Level: "lossless",
	}
	code, resp, _ := service.SongUrl()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
