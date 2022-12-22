package service

import (
	"fmt"
	"testing"
)

func TestPlaylistDetailService_PlaylistDetail(t *testing.T) {
	service := &PlaylistTrackAllService{
		Id: "139746382",
	}
	_, resp := service.AllTracks()
	fmt.Println(string(resp))
}
