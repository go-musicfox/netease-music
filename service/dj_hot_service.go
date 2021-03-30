package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjHotService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *DjHotService) DjHot() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {

		data["offset"] = service.Offset
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/hot/v1`, data, options)

	return reBody
}
