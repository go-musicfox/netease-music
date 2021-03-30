package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type EventForwardService struct {
	Uid      string `json:"uid" form:"uid"`
	EvId     string `json:"evId" form:"evId"`
	Forwards string `json:"forwards" form:"forwards"`
}

func (service *EventForwardService) EventForward() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.EvId
	data["eventUserId"] = service.Uid
	data["forwards"] = service.Forwards
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/event/forward`, data, options)

	return reBody
}
