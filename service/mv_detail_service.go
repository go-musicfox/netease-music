package service

import (
	"github.com/anhoder/netease-music/util"
)

type MvDetailService struct {
	ID string `json:"mvid" form:"mvid"`
}

func (service *MvDetailService) MvDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v1/mv/detail`, data, options)

	return reBody
}
