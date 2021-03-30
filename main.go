package main

import (
    "fmt"
    "github.com/anhoder/netease-music/service"
    "github.com/anhoder/netease-music/util"
    "github.com/telanflow/cookiejar"
)

func main() {
    cookieJar, _ := cookiejar.NewFileJar("cookie.txt", nil)
    util.SetGlobalCookieJar(cookieJar)
    songs := service.RecommendSongsService{}
    fmt.Println(songs.RecommendSongs());
}
