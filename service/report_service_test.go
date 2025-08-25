package service

import (
	"testing"

	"github.com/go-musicfox/netease-music/util"
)

func TestReportService(t *testing.T) {
	// 为确保测试成功，请将登陆后的cookies以map类型传入
	cookiesMap := map[string]string{}
	// 添加 cookies 到全局 CookieJar
	util.AddCookiesToJar(util.GetGlobalCookieJar(), cookiesMap, "https://music.163.com")

	service := &ReportService{
		ID:         2084034562,
		Type:       "song",
		SourceType: "list",
		SourceId:   "627357921",
		Time:       30000,
		EndType:    "playend",
	}

	code, bodyBytes, err := service.Playend()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	t.Logf("code: %f, body: %s", code, string(bodyBytes))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
	if len(bodyBytes) == 0 {
		t.Error("response body is empty")
	}
}
