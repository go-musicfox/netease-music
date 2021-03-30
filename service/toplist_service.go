package service

import (
	"github.com/anhoder/netease-music/util"
)

type ToplistService struct {
}

func (service *ToplistService) Toplist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "linuxapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/toplist`, data, options)

	return reBody
}
