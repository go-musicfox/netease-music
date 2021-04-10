package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistHotService struct{}

func (service *PlaylistHotService) PlaylistHot() (float64, string) {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/hottags`, data, options)

	return code, reBody
}
