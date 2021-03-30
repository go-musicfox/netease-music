package service

import (
	"github.com/anhoder/netease-music/util"
)

type AlbumDetailService struct {
	ID string `json:"id" form:"id"`
}

func (service *AlbumDetailService) AlbumDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/vipmall/albumproduct/detail`, data, options)

	return reBody
}
