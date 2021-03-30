package service

import (
	"github.com/anhoder/netease-music/util"
)

type AlbumSublistService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *AlbumSublistService) AlbumSublist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "25"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["total"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/album/sublist`, data, options)

	return reBody
}
