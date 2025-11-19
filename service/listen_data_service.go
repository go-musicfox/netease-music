package service

import (
	"encoding/json"

	"github.com/go-musicfox/netease-music/util"
)

type ListenDataService struct {
}

// ListenDataTodaySongsResult 今日听过的歌曲结果
type ListenDataTodaySongsResult struct {
	Code float64 `json:"code"`
	Data struct {
		SongDTOs songDTOs `json:"songDTOs"`
	} `json:"data"`
	Message string `json:"message"`
}

type songDTOs struct {
	SongId       int     `json:"songId"`
	SongName     string  `json:"songName"`
	AliasName    string  `json:"aliasName"`
	LastPlayTime float64 `json:"lastPlayTime"`
}

// ListenDataTotalResult 听歌总时长结果
type ListenDataTotalResult struct {
	Code float64 `json:"code"`
	Data struct {
		TotalDuration float64 `json:"totalDuration"`
	} `json:"data"`
	Message string `json:"message"`
}

// TodaySongs 今日听过的歌曲
func (service *ListenDataService) TodaySongs() (result ListenDataTodaySongsResult, err error) {
	data := map[string]interface{}{}
	api := "https://music.163.com/api/content/activity/listen/data/today/song/play/rank"
	_, bytesData, err := util.CallWeapi(api, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytesData, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}

// Total 听歌总时长
func (service *ListenDataService) Total() (result ListenDataTotalResult, err error) {
	data := map[string]interface{}{}
	api := "https://music.163.com/api/content/activity/listen/data/total"
	_, bytesData, err := util.CallWeapi(api, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytesData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
