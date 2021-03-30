package service

import (
	"github.com/anhoder/netease-music/util"
)

type VideoGroupListService struct {
}

func (service *VideoGroupListService) VideoGroupList() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/cloudvideo/group/list`, data, options)

	return reBody
}
