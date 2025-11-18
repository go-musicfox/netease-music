package service

import (
	"github.com/go-musicfox/netease-music/util"
)

type YunbeiSigninService struct {
}

func (service *YunbeiSigninService) Signin() (float64, []byte, error) {
	cookiejar := util.GetGlobalCookieJar()
	csrfToken := util.GetCsrfToken(cookiejar)
	data := map[string]interface{}{
		"type": "1",
		"checkToken": "", // 这个值由易盾通过浏览器指纹动态生成
		"csrf_token": csrfToken,
	}
	api := "https://music.163.com/weapi/point/dailyTask?csrf_token=" + csrfToken
	code, bytesData, err := util.CallWeapi(api, data)
	return code, bytesData, err
}
