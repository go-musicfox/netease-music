package service

import (
	"github.com/go-musicfox/netease-music/util"
)

type YunbeiSigninService struct {
}

func (service *YunbeiSigninService) Signin() (float64, []byte, error) {
	data := map[string]interface{}{
		"type": "0",
	}
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	api := "https://music.163.com/weapi/point/dailyTask?csrf_token=" + csrfToken
	code, bytesData, err := util.CallWeapi(api, data)
	return code, bytesData, err
}
