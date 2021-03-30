package service

import (
	"github.com/anhoder/netease-music/util"
)

type FmTrashService struct {
	SongID string `json:"id" form:"id"`
}

func (service *FmTrashService) FmTrash() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["songId"] = service.SongID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/radio/trash/add?alg=RT&songId=`+service.SongID+`&time=25`, data, options)

	return reBody
}
