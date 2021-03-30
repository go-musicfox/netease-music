package service

import (
	"github.com/anhoder/netease-music/util"
)

type PersonalizedPrivatecontentService struct {
}

func (service *PersonalizedPrivatecontentService) PersonalizedPrivatecontent() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/privatecontent`, data, options)

	return reBody
}
