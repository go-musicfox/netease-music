package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type RecommendSongsService struct {
	ID string `json:"id" form:"id"`
}

func (service *RecommendSongsService) RecommendSongs() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "ios"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v3/discovery/recommend/songs`, data, options)

	return reBody
}
