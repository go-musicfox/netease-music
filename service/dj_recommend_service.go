package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjRecommendService struct {
}

func (service *DjRecommendService) DjRecommend() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/recommend/v1`, data, options)

	return reBody
}
