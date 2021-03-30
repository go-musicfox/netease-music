package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type LyricService struct {
	ID string `json:"id" form:"id"`
}

func (service *LyricService) Lyric() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "linuxapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["lv"] = "-1"
	data["kv"] = "-1"
	data["tv"] = "-1"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/song/lyric`, data, options)

	return reBody
}
