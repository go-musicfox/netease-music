package service

import (
	"github.com/anhoder/netease-music/util"
)

type TopArtistsService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *TopArtistsService) TopArtists() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	if service.Limit == "" {
		data["limit"] = "50"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/top`, data, options)

	return reBody
}
