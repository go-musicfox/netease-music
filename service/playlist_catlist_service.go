package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistCatlistService struct{}

func (service *PlaylistCatlistService) PlaylistCatlist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/catalogue`, data, options)

	return reBody
}
