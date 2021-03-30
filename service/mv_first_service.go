package service

import (
	"github.com/anhoder/netease-music/util"
)

type MvFirstService struct {
	Area  string `json:"area" form:"area"`
	Limit string `json:"limit" form:"limit"`
}

func (service *MvFirstService) MvFirst() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["area"] = service.Area
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}

	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://interface.music.163.com/weapi/mv/first`, data, options)

	return reBody
}
