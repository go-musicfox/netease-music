package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjCategoryRecommendService struct {
}

func (service *DjCategoryRecommendService) DjCategoryRecommend() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/home/category/recommend`, data, options)

	return reBody
}
