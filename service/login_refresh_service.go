package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type LoginRefreshService struct {
}

func (service *LoginRefreshService) LoginRefresh() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/login/token/refresh`, data, options)

	return reBody
}
