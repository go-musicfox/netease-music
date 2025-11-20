package service

import (
	"testing"

	"github.com/go-musicfox/netease-music/util"
)

func TestListenData(t *testing.T) {
	cookieMap := util.ParseCookieString("")
	util.AddCookiesToJar(util.GetGlobalCookieJar(), cookieMap, "https://music.163.com")
	service := &ListenDataService{}
	result, err := service.Total()
	if err != nil {
		t.Logf("error: %v", err)
	}
	if result.Code != 200 {
		t.Logf("发生错误，code：%f", result.Code)
	}
	t.Logf("总计听歌时长: %.2f 秒", result.Data.TotalDuration)
	report, err := service.GetReport("week")
	if err != nil {
		t.Logf("error: %v", err)
	}
	if result.Code != 200 {
		t.Logf("发生错误，code：%f", result.Code)
	}
	t.Log("\n" + Format(report))
}
