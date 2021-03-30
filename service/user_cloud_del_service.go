package service

import (
	"github.com/anhoder/netease-music/util"
)

type UserCloudDelService struct {
	ID string `json:"id" form:"id"`
}

func (service *UserCloudDelService) UserCloudDel() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["songIds"] = "[" + service.ID + "]"
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/cloud/del`, data, options)

	return reBody
}
