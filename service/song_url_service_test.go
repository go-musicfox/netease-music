package service

import (
	"fmt"
	"testing"

	"github.com/anhoder/netease-music/util"
)

func TestSongUrlService_SongUrl(t *testing.T) {
	service := &SongUrlService{
		ID: "1962165890",
		Br: "999000",
	}
	code, resp := service.SongUrl()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}

func TestSongUrlService_SongUrlWithUNM(t *testing.T) {
	util.UNMSwitch = true
	util.Sources = []string{"kuwo"}
	service := &SongUrlService{
		ID: "1962165890",
		Br: "320000",
	}
	code, resp := service.SongUrl()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
