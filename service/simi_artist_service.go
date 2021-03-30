package service

import (
	"github.com/anhoder/netease-music/util"
)

type SimiArtistService struct {
	ID string `json:"id" form:"id"`
}

func (service *SimiArtistService) SimiArtist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/discovery/simiArtist`, data, options)

	return reBody
}
