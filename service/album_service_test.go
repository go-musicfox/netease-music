package service

import (
	"fmt"
	"testing"
)

func TestAlbumService_Album(t *testing.T) {
	service := &AlbumService{
		ID: "147779282",
	}
	code, resp := service.Album()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
