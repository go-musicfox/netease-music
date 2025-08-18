package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"

	"github.com/go-musicfox/netease-music/util"
)

type LoginCellphoneService struct {
	Phone       string `json:"phone" form:"phone"`
	Countrycode string `json:"countrycode" form:"countrycode"`
	Password    string `json:"password" form:"password"`
	Md5password string `json:"md5_password" form:"md5_password"`
	Captcha     string `json:"captcha" from:"captcha"`
	CsrfToken   string `json:"csrf_token" from:"csrf_token"`
}

func (service *LoginCellphoneService) LoginCellphone() (float64, []byte) {

	// cookiesOS := &http.Cookie{Name: "os", Value: "ios"}
	// appVersion := &http.Cookie{Name: "appver", Value: "8.7.01"}

	// options := &util.Options{
	// 	Crypto:  "weapi",
	// 	Ua:      "pc",
	// 	Cookies: []*http.Cookie{cookiesOS, appVersion},
	// }
	data := make(map[string]string)

	data["phone"] = service.Phone
	if service.Countrycode != "" {
		data["countrycode"] = service.Countrycode
	} else {
		data["countrycode"] = "86"
	}

	if service.Captcha != "" {
		data["captcha"] = service.Captcha
	}

	data["csrf_token"] = service.CsrfToken

	if service.Password != "" {
		h := md5.New()
		h.Write([]byte(service.Password))
		data["password"] = hex.EncodeToString(h.Sum(nil))
	} else {
		data["password"] = service.Md5password
	}
	data["rememberLogin"] = "true"

	// code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/login/cellphone`, data, options)

	api := "https://music.163.com/weapi/login/cellphone"
	req := util.NewRequest(api)
	// 传入加密后的formdata
	req.Datas = util.Weapi(data)
	resp := req.SendPost()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}
	defer resp.Body.Close()
	// 获取api调用后的code
	var respData map[string]interface{}
	err = json.Unmarshal(bodyBytes, &respData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	code, ok := respData["code"].(float64)
	if !ok {
		log.Fatal("Could not get 'code' or it's not a number")
	}
	return code, bodyBytes
}
