# netease-music

**fork自 [sirodeneko/NeteaseCloudMusicApiWithGo](https://github.com/sirodeneko/NeteaseCloudMusicApiWithGo)**

在原项目的基础上去除API功能，只保留service、util作为一个独立的包，方便在其他go项目中调用。

## 功能特性
1. 登录
2. 刷新登录
3. 发送验证码
4. 校验验证码
5. 注册(修改密码)  
6. 。。。等160多个api
7. 支持UNM

## 环境要求

文档参考[NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)

需要 Golang 1.12以上 环境

## 运行

```shell
go get -u github.com/go-musicfox/netease-music
```

```go
package main

import (
    "fmt"
    "github.com/go-musicfox/netease-music/service"
    "github.com/go-musicfox/netease-music/util"
    "github.com/telanflow/cookiejar"
)

func main() {
    cookieJar, _ := cookiejar.NewFileJar("cookie.txt", nil)
    util.SetGlobalCookieJar(cookieJar)
    songs := service.RecommendSongsService{}
    fmt.Println(songs.RecommendSongs());
}

```


