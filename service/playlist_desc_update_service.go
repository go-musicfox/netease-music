package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistDescUpdateService struct {
	Id   string `json:"id" form:"id"`
	Desc string `json:"desc" form:"desc"`
}

func (service *PlaylistDescUpdateService) PlaylistDescUpdate() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/playlist/desc/update",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["desc"] = service.Desc
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/desc/update`, data, options)

	return reBody
}
