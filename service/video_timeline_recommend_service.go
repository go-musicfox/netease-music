package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type VideoTimelineRecommendService struct {
	Offset string `json:"offset" form:"offset"`
}

func (service *VideoTimelineRecommendService) VideoTimelineRecommend() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["filterLives"] = "[]"
	data["withProgramInfo"] = "true"
	data["needUrl"] = "1"
	data["resolution"] = "480"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/videotimeline/get`, data, options)

	return reBody
}
