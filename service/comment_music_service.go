package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type CommentMusicService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Before string `json:"before" form:"before"`
}

func (service *CommentMusicService) CommentMusic() (float64, []byte) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["rid"] = service.ID
	if service.Limit == "" {
		data["limit"] = "20"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.Before == "" {
		data["beforeTime"] = "0"
	} else {
		data["beforeTime"] = service.Before
	}
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v1/resource/comments/R_SO_4_`+service.ID, data, options)

	return code, reBody
}
