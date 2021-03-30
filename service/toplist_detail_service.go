package service

import (
	"github.com/anhoder/netease-music/util"
)

type ToplistDetailService struct {
}

func (service *ToplistDetailService) ToplistDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/toplist/detail`, data, options)

	return reBody
}
