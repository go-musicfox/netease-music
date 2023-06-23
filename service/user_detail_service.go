package service

import (
	"github.com/go-musicfox/netease-music/util"
)

type UserDetailService struct {
	Uid string `json:"uid" form:"uid"`
}

func (service *UserDetailService) UserDetail() (float64, []byte) {

	options := &util.Options{
		Crypto: "weapi",
	}
	data := make(map[string]string)
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/user/detail/`+service.Uid, data, options)

	return code, reBody
}
