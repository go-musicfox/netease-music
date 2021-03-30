package service

import (
	"github.com/anhoder/netease-music/util"
)

type UserRecordService struct {
	UId  string `json:"uid" form:"uid"`
	Type string `json:"type" form:"type"`
}

func (service *UserRecordService) UserRecord() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["uid"] = service.UId

	if service.Type == "1" {
		data["type"] = "1"
	} else {
		data["type"] = "0"
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/play/record`, data, options)

	return reBody
}
