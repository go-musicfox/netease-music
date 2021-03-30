package service

import (
	"github.com/anhoder/netease-music/util"
)

type PersonalizedNewsongService struct {
}

func (service *PersonalizedNewsongService) PersonalizedNewsong() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	data["type"] = "recommend"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/newsong`, data, options)

	return reBody
}
