package service

import (
	"github.com/anhoder/netease-music/util"
)

type SettingService struct {
}

func (service *SettingService) Setting() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/user/setting`, data, options)

	return reBody
}
