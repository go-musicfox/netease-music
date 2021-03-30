package service

import (
	"github.com/anhoder/netease-music/util"
)

type RebindService struct {
	Captcha    string `json:"captcha" form:"captcha"`
	Phone      string `json:"phone" form:"phone"`
	Oldcaptcha string `json:"oldcaptcha" form:"oldcaptcha"`
	Ctcode     string `json:"ctcode" form:"ctcode"`
}

func (service *RebindService) Rebind() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	data["phone"] = service.Phone
	data["captcha"] = service.Captcha
	data["captcha"] = service.Captcha
	data["oldcaptcha"] = service.Oldcaptcha

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/user/replaceCellphone`, data, options)

	return reBody
}