package service

import (
	"github.com/anhoder/netease-music/util"
)

type SearchMultimatchService struct {
	Type string `json:"type" form:"type"`
	S    string `json:"keywords" form:"keywords"`
}

func (service *SearchMultimatchService) SearchMultimatch() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	if service.Type == "" {
		service.Type = "1"
	}
	data["type"] = service.Type
	data["s"] = service.S
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/search/suggest/multimatch`, data, options)

	return reBody
}
