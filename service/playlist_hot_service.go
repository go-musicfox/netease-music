package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistHotService struct{}

func (service *PlaylistHotService) PlaylistHot() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/hottags`, data, options)

	return reBody
}
