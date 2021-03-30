package service

import (
	"github.com/anhoder/netease-music/util"
)

type ActivateInitProfileService struct {
	Nickname string `json:"nickname" form:"nickname"`
}

func (service *ActivateInitProfileService) ActivateInitProfile() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/activate/initProfile",
	}
	data := make(map[string]string)
	data["nickname"] = service.Nickname

	reBody, _ := util.CreateRequest("POST", `http://music.163.com/eapi/activate/initProfile`, data, options)

	return reBody
}
