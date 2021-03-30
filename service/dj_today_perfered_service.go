package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjTodayPerferedService struct {
	Page string `json:"page" form:"page"`
}

func (service *DjTodayPerferedService) DjTodayPerfered() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	if service.Page == "" {
		data["page"] = "0"
	} else {
		data["page"] = service.Page
	}
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/home/today/perfered`, data, options)

	return reBody
}