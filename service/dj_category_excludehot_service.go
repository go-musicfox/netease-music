package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjCategoryExcludehotService struct {
}

func (service *DjCategoryExcludehotService) DjCategoryExcludehot() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/category/excludehot`, data, options)

	return reBody
}
