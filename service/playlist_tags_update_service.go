package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistTagsUpdateService struct {
	Id   string `json:"id" form:"id"`
	Tags string `json:"tags" form:"tags"`
}

func (service *PlaylistTagsUpdateService) PlaylistTagsUpdate() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/playlist/tags/update",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["tags"] = service.Tags
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/tags/update`, data, options)

	return reBody
}
