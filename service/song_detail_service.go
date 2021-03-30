package service

import (
	"encoding/json"
	"net/http"
	"github.com/anhoder/netease-music/util"
	"strings"
)

type SongDetailService struct {
	Ids string `json:"ids" form:"ids"`
}

func (service *SongDetailService) SongDetail() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}

	type IDS struct {
		ID string `json:"id"`
	}

	var cids []IDS

	strs := strings.Split(service.Ids, ",")
	for _, item := range strs {
		cids = append(cids, IDS{ID: item})
	}
	sidsJsonByte, _ := json.Marshal(cids)

	data := make(map[string]string)
	data["c"] = string(sidsJsonByte)
	data["ids"] = "[" + service.Ids + "]"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v3/song/detail`, data, options)

	return reBody
}
