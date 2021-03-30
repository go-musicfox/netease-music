package service

import (
	"github.com/anhoder/netease-music/util"
)

type VideoDetailInfoService struct {
	ID string `json:"vid" form:"vid"`
}

func (service *VideoDetailInfoService) VideoDetailInfo() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["threadid"] = "R_VI_62_" + service.ID
	data["composeliked"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/comment/commentthread/info`, data, options)

	return reBody
}
