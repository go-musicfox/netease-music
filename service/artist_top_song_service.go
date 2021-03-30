package service

import (
	"github.com/anhoder/netease-music/util"
)

type ArtistTopSongService struct {
	Id string `json:"id" form:"id"`
}

func (service *ArtistTopSongService) ArtistTopSong() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	data["id"] = service.Id

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/artist/top/song`, data, options)

	return reBody
}
