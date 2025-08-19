package service

import (
	"encoding/json"
	"strconv"

	"github.com/go-musicfox/netease-music/util"
)

type ReportService struct {
	ID         int64  `json:"id" form:"id"`
	Type       string `json:"type" form:"type"`             // song:歌曲, dj:播客
	SourceType string `json:"sourceType" form:"sourceType"` // list, album, dailySongRecommend, userfm
	SourceId   string `json:"sourceId" form:"sourceId"`
	Time       int64  `json:"time" form:"time"`
	Alg        string `json:"alg" form:"alg"`
	EndType    string `json:"endType" form:"endType"` // playend：正常结束；interrupt：第三方APP打断： exception: 错误； ui: 用户切歌
}

func (service *ReportService) Playend() (float64, []byte) {
	if service.EndType == "" {
		service.EndType = "playend"
	}
	if service.Type == "" {
		service.Type = "song"
	}

	jsonData := map[string]interface{}{
		"type":        service.Type,
		"wifi":        0,
		"download":    0,
		"id":          service.ID,
		"time":        service.Time,
		"end":         service.EndType,
		"source":      service.SourceType,
		"mainsite":    "1",
		"mainsiteWeb": "1",
		"content":     "",
	}

	if service.SourceId != "" {
		jsonData["sourceId"] = service.SourceId
		if _, err := strconv.ParseInt(service.SourceId, 10, 64); err == nil {
			jsonData["content"] = "id=" + service.SourceId
		}
	}

	if service.Alg != "" {
		jsonData["alg"] = service.Alg
	}

	logs := []map[string]interface{}{
		{
			"action": "play",
			"json":   jsonData,
		},
	}

	data := make(map[string]string)
	if str, err := json.Marshal(logs); err == nil {
		data["logs"] = string(str)
	}

	api := "https://clientlogusf.music.163.com/weapi/feedback/weblog"
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	data["csrf_token"] = csrfToken
	code, bodyBytes := util.CallWeapi(api+"?csrf_token="+csrfToken, data)

	return code, bodyBytes
}

func (service *ReportService) Playstart() (float64, []byte) {
	if service.Type == "" {
		service.Type = "song"
	}

	jsonData := map[string]interface{}{
		"id":          service.ID,
		"type":        service.Type,
		"content":     "",
		"mainsite":    "1",
		"mainsiteWeb": "1",
	}

	if _, err := strconv.ParseInt(service.SourceId, 10, 64); err == nil {
		jsonData["content"] = "id=" + service.SourceId
	}

	if service.Alg != "" {
		jsonData["alg"] = service.Alg
	}

	logs := []map[string]interface{}{
		{
			"action": "startplay",
			"json":   jsonData,
		},
	}

	data := make(map[string]string)
	if str, err := json.Marshal(logs); err == nil {
		data["logs"] = string(str)
	}
	api := "https://clientlogusf.music.163.com/weapi/feedback/weblog"
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	data["csrf_token"] = csrfToken
	code, bodyBytes := util.CallWeapi(api+"?csrf_token="+csrfToken, data)

	return code, bodyBytes
}

// 获取P2P流量开关状态
func (service *ReportService) P2PFlowSwitchGet() (float64, []byte) {
	data := make(map[string]string)
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	data["csrf_token"] = csrfToken
	api := "https://music.163.com/weapi/activity/p2p/flow/switch/get"
	code, bodyBytes := util.CallWeapi(api+"?csrf_token="+csrfToken, data)
	return code, bodyBytes
}
