package service

import (
	"github.com/anhoder/netease-music/util"
)

type PersonalizedMvService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *PersonalizedMvService) PersonalizedMv() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/mv`, data, options)

	return reBody
}
