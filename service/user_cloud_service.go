package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type UserCloudService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *UserCloudService) UserCloud() map[string]interface{} {

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
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/cloud/get`, data, options)

	return reBody
}
