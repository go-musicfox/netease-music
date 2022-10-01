module github.com/anhoder/netease-music

go 1.12

require (
	github.com/asmcos/requests v0.0.0-20200725080539-508ff1c69720
	github.com/buger/jsonparser v1.1.1
	github.com/cnsilvan/UnblockNeteaseMusic v0.0.0-20220606141015-ee827820fabb
	github.com/forgoer/openssl v0.0.0-20200331032942-ad9f8d57d8b1
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	github.com/asmcos/requests => ./staging/github.com/asmcos/requests
)
