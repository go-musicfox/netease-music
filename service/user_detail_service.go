package service

import (
	"github.com/anhoder/netease-music/util"
)

type UserDetailService struct {
	Uid string `json:"uid" form:"uid"`
}

func (service *UserDetailService) UserDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/user/detail/`+service.Uid, data, options)

	return reBody
}
