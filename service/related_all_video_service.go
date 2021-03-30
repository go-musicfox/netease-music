package service

import (
	"github.com/anhoder/netease-music/util"
)

type RelatedAllVideoService struct {
	ID string `json:"id" form:"id"`
}

func (service *RelatedAllVideoService) RelatedAllVideo() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["type"] = "1"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/cloudvideo/v1/allvideo/rcmd`, data, options)

	return reBody
}
