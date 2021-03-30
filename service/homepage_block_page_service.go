package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type HomepageBlockPageService struct {
	Refresh string `json:"refresh" form:"refresh"`
}

func (service *HomepageBlockPageService) HomepageBlockPage() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Refresh == "" {
		service.Refresh = "true"
	}
	data["refresh"] = service.Refresh
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/homepage/block/page`, data, options)

	return reBody
}
