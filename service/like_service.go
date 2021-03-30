package service

import (
	"github.com/anhoder/netease-music/util"
)

type LikeService struct {
	ID string `json:"id" form:"id"`
	L  string `json:"like" form:"like"`
}

func (service *LikeService) Like() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["trackId"] = service.ID
	if service.L == "" {
		data["like"] = "true"
	} else {
		data["like"] = service.L
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/radio/like?alg=itembased&trackId=`+service.ID+`&time=25`, data, options)

	return reBody
}
