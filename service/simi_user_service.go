package service

import (
	"github.com/anhoder/netease-music/util"
)

type SimiUserService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *SimiUserService) SimiUser() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["songid"] = service.ID
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

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/discovery/simiUser`, data, options)

	return reBody
}
