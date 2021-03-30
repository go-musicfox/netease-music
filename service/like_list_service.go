package service

import (
	"github.com/anhoder/netease-music/util"
)

type LikeListService struct {
	UID string `json:"uid" form:"uid"`
}

func (service *LikeListService) LikeList() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}

	data := make(map[string]string)
	data["uid"] = service.UID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/song/like/get`, data, options)

	return reBody
}
