package service

import (
	"github.com/anhoder/netease-music/util"
)

type PersonalFmService struct {
}

func (service *PersonalFmService) PersonalFm() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/radio/get`, data, options)

	return reBody
}
