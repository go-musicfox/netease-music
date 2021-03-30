package service

import (
	"github.com/anhoder/netease-music/util"
)

type UserAudioService struct {
	UID string `json:"uid" form:"uid"`
}

func (service *UserAudioService) UserAudio() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["userId"] = service.UID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/get/byuser`, data, options)

	return reBody
}
