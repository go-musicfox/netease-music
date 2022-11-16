package service

import (
	"net/http"

	"github.com/anhoder/netease-music/util"
)

type SongUrlV1Service struct {
	ID         string `json:"id" form:"id"`
	Level      string `json:"level" form:"level"` // standard,higher,exhigh,lossless,hires
	EncodeType string `json:"encodeType" form:"encodeType"`
}

func (service *SongUrlV1Service) SongUrl() (float64, []byte) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: []*http.Cookie{cookiesOS},
		Url:     "/api/song/enhance/player/url/v1",
	}
	data := make(map[string]string)
	data["ids"] = "[" + service.ID + "]"
	if service.Level == "" {
		service.Level = "higher"
	}
	data["level"] = service.Level
	data["encodeType"] = service.EncodeType

	code, reBody, _ := util.CreateRequest("POST", `https://interface.music.163.com/eapi/song/enhance/player/url/v1`, data, options)

	return code, reBody
}
