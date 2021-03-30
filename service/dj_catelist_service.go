package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjCatelistService struct {
}

func (service *DjCatelistService) DjCatelist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/category/get`, data, options)

	return reBody
}
