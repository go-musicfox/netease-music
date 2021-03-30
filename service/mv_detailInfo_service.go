package service

import (
	"github.com/anhoder/netease-music/util"
)

type MvDetailInfoService struct {
	ID string `json:"mvid" form:"mvid"`
}

func (service *MvDetailInfoService) MvDetailInfo() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["threadid"] = "R_MV_5_" + service.ID
	data["composeliked"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/comment/commentthread/info`, data, options)

	return reBody
}