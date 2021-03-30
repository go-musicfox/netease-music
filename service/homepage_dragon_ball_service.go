package service

import (
	"github.com/anhoder/netease-music/util"
)

type HomepageDragonBallService struct {
}

func (service *HomepageDragonBallService) HomepageDragonBall() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/homepage/dragon/ball/static",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/eapi/homepage/dragon/ball/static`, data, options)

	return reBody
}
