package service

import (
	"github.com/anhoder/netease-music/util"
)

type SearchSuggestService struct {
	S    string `json:"keywords" form:"keywords"`
	Type string `json:"type" form:"type"`
}

func (service *SearchSuggestService) SearchSuggest() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	if service.Type == "mobile" {
		service.Type = "keyword"
	} else {
		service.Type = "web"
	}
	data["s"] = service.S
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/search/suggest/`+service.Type, data, options)

	return reBody
}