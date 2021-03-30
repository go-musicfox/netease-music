package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjDetailService struct {
	ID string `json:"rid" form:"rid"`
}

func (service *DjDetailService) DjDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/get`, data, options)

	return reBody
}
