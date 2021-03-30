package service

import (
	"github.com/anhoder/netease-music/util"
)

type AlbumSubService struct {
	ID string `json:"id" form:"id"`
	T  string `json:"t" form:"t"`
}

func (service *AlbumSubService) AlbumSub() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	if service.T == "1" {
		service.T = "sub"
	} else {
		service.T = "unsub"
	}
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/album/`+service.T, data, options)

	return reBody
}
