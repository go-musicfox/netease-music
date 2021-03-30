package service

import (
	"github.com/anhoder/netease-music/util"
)

type VideoGroupService struct {
	GroupID string `json:"id" form:"id"`
	Offset  string `json:"offset" form:"offset"`
}

func (service *VideoGroupService) VideoGroup() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["groupId"] = service.GroupID
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	data["need_preview_url"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/videotimeline/videogroup/otherclient/get`, data, options)

	return reBody
}
