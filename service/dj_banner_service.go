package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type DjBannerService struct {
}

func (service *DjBannerService) DjBanner() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/banner/get`, data, options)

	return reBody
}
