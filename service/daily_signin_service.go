package service

import (
	"github.com/anhoder/netease-music/util"
)

type DailySigninService struct {
	Type string `json:"type" form:"type"`
}

func (service *DailySigninService) DailySignin() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	if service.Type == "" {
		data["type"] = "0"
	} else {
		data["type"] = service.Type
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/point/dailyTask`, data, options)

	return reBody
}