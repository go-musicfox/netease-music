package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type PersonalizedService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *PersonalizedService) Personalized() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	data["order"] = "true"
	data["n"] = "1000"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/playlist`, data, options)

	return reBody
}
