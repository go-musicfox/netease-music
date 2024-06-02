module github.com/go-musicfox/netease-music

go 1.18

require (
	github.com/go-musicfox/requests v0.2.2
	github.com/buger/jsonparser v1.1.1
	github.com/cnsilvan/UnblockNeteaseMusic v0.0.0-20230310083816-92b59c95a366
	github.com/forgoer/openssl v1.6.0
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/text v0.10.0 // indirect
)

replace (
	github.com/cnsilvan/UnblockNeteaseMusic => github.com/go-musicfox/UnblockNeteaseMusic v0.1.3
)
