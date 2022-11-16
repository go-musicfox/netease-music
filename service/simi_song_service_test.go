package service

import (
	"fmt"
	"testing"
)

func TestSimiSongService_SimiSong(t *testing.T) {
	service := &SimiSongService{
		ID:    "405998841",
		Limit: "999",
	}
	code, resp := service.SimiSong()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
