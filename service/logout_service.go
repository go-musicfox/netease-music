package service

import (
	"github.com/go-musicfox/netease-music/util"
)

type LogoutService struct {
}

func (service *LogoutService) Logout() (float64, []byte) {
	api := "https://music.163.com/weapi/logout"
	data := make(map[string]interface{})
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	data["csrf_token"] = csrfToken
	code, bodyBytes := util.CallWeapi(api, data)
	return code, bodyBytes
}
