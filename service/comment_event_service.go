package service

import (
	"github.com/anhoder/netease-music/util"
)

type CommentEventService struct {
	ThreadId   string `json:"threadId" form:"threadId"`
	Limit      string `json:"limit" form:"limit"`
	Offset     string `json:"offset" form:"offset"`
	BeforeTime string `json:"beforeTime" form:"beforeTime"`
}

func (service *CommentEventService) CommentEvent() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "20"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.BeforeTime == "" {
		data["beforeTime"] = "0"
	} else {
		data["beforeTime"] = service.BeforeTime
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/resource/comments/`+service.ThreadId, data, options)

	return reBody
}
