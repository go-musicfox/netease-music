package service

import (
	"github.com/anhoder/netease-music/util"
)

type UserSubcountService struct {
}

func (service *UserSubcountService) UserSubcount() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/subcount`, data, options)

	return reBody
}
