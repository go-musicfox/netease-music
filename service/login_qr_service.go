package service

import (
	"encoding/json"
	"log"

	"github.com/go-musicfox/netease-music/util"
)

type LoginQRService struct {
	UniKey string `json:"unikey"`
}

func (service *LoginQRService) GetKey() (float64, []byte, string) {
	data := map[string]interface{}{
		"type":         1,
		"noCheckToken": true,
	}

	api := "https://music.163.com/weapi/login/qrcode/unikey"
	code, bodyBytes := util.CallWeapi(api, data)
	if code != 200 || len(bodyBytes) == 0 {
		return code, bodyBytes, ""
	}
	err := json.Unmarshal(bodyBytes, service)
	if err != nil {
		log.Fatalf("Error unmarshalling bodybytes: %v", err)
	}

	// 生成 chainId，这个是新版本新加的参数
	cookieJar := util.GetGlobalCookieJar()
	chainID := util.GenerateChainID(cookieJar)
	qrcodeUrl := ("http://music.163.com/login?codekey=" +
		service.UniKey + "&chainId=" + chainID)
	return code, bodyBytes, qrcodeUrl
}

func (service *LoginQRService) CheckQR() (float64, []byte) {
	if service.UniKey == "" {
		return 0, nil
	}
	data := map[string]interface{}{
		"type":         1,
		"noCheckToken": true,
		"key":          service.UniKey,
	}

	api := "https://music.163.com/weapi/login/qrcode/client/login"
	code, bodyBytes := util.CallWeapi(api, data)
	return code, bodyBytes
}
