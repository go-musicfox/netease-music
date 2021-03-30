package service

import (
	"github.com/anhoder/netease-music/util"
)

type CountriesCodeListService struct {
}

func (service *CountriesCodeListService) CountriesCodeList() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/lbs/countries/v1",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/lbs/countries/v1`, data, options)

	return reBody
}