package service

import (
	"fmt"
	"testing"

	"github.com/anhoder/netease-music/util"
)

func TestAlbumService_Album(t *testing.T) {
	util.UNMSwitch = true
	util.Sources = []string{"kuwo"}
	service := &AlbumService{
		ID: "147779282",
	}
	code, resp := service.Album()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
