package service

import (
	"net/http"

	"github.com/go-musicfox/netease-music/util"
)

type ArtistsSubListService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *ArtistsSubListService) ArtistsSubList() (float64, []byte) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)

	if service.Limit == "" {
		service.Limit = "30"
	}
	if service.Offset == "" {
		service.Offset = "0"
	}
	data["limit"] = service.Limit
	data["offset"] = service.Offset
	data["total"] = "true"

	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/sublist`, data, options)

	return code, reBody
}
