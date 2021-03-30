package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type PlaylistOrderUpdateService struct {
	Ids string `json:"ids" form:"ids"`
}

func (service *PlaylistOrderUpdateService) PlaylistOrderUpdate() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.Ids
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/playlist/order/update`, data, options)

	return reBody
}
