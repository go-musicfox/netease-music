package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type PlaylistCreateService struct {
	Name    string `json:"name" form:"name"`
	Privacy string `json:"privacy" form:"privacy"`
}

func (service *PlaylistCreateService) PlaylistCreate() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Privacy != "10" {
		service.Privacy = "0"
	}
	data["name"] = service.Name
	data["privacy"] = service.Privacy
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/create`, data, options)

	return reBody
}
