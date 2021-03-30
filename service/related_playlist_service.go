package service

import (
	"regexp"
	"github.com/anhoder/netease-music/util"
)

type RelatedPlaylistService struct {
	ID string `json:"id" form:"id"`
}

func (service *RelatedPlaylistService) RelatedPlaylist() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
		Ua:      "pc",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("GET", `https://music.163.com/playlist?id=`+service.ID, data, options)

	reg := regexp.MustCompile("<div class=\"cver u-cover u-cover-3\">[\\s\\S]*?<img src=\"([^\"]+)\">[\\s\\S]*?<a class=\"sname f-fs1 s-fc0\" href=\"([^\"]+)\"[^>]*>([^<]+?)<\\/a>[\\s\\S]*?<a class=\"nm nm f-thide s-fc3\" href=\"([^\"]+)\"[^>]*>([^<]+?)<\\/a>")
	results := reg.FindAllStringSubmatch(reBody["html"].(string), -1)

	type Creator struct {
		UserId   string `json:"userId"`
		Nickname string `json:"nickname"`
	}

	type Result struct {
		Id          string  `json:"id"`
		Name        string  `json:"name"`
		CoverImgUrl string  `json:"coverImgUrl"`
		Creator     Creator `json:"creator"`
	}
	var Results []Result
	for _, result := range results {
		var item Result
		item.Id = result[2][len("/playlist?id="):]
		item.Name = result[3]
		item.CoverImgUrl = result[1][0 : len(result[1])-len("?param=50y50")]
		item.Creator.UserId = result[4][len("/user/home?id="):]
		item.Creator.Nickname = result[5]
		Results = append(Results, item)
	}

	delete(reBody, "html")
	reBody["playlists"] = Results
	return reBody
}
