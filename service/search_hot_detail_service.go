package service

import (
	"github.com/anhoder/netease-music/util"
)

type SearchHotDetailService struct {
}

func (service *SearchHotDetailService) SearchHotDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/hotsearchlist/get`, data, options)

	return reBody
}
