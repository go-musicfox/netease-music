package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjCatelistService struct {
}

func (service *DjCatelistService) DjCatelist() (float64, string) {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/category/get`, data, options)

	return code, reBody
}
