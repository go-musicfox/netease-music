package service

import (
	"github.com/anhoder/netease-music/util"
)

type SearchHotService struct {
}

func (service *SearchHotService) SearchHot() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
		Ua:      "mobile",
	}
	data := make(map[string]string)
	data["type"] = "1111"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/search/hot`, data, options)

	return reBody
}
