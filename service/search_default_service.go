package service

import (
	"github.com/anhoder/netease-music/util"
)

type SearchDefaultService struct {
}

func (service *SearchDefaultService) SearchDefault() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/search/defaultkeyword/get",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/search/defaultkeyword/get`, data, options)

	return reBody
}
