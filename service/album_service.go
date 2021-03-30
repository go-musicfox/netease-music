package service

import (
	"github.com/anhoder/netease-music/util"
)

type AlbumService struct {
	ID string `json:"id" form:"id"`
}

func (service *AlbumService) Album() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/album/`+service.ID, data, options)

	return reBody
}
