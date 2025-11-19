package service

import "encoding/json"

type BaseServiceResult struct {
}

// FormatStruct 将结构体格式化为json字符串形式
func FormatStruct(v interface{}) string {
	bytesData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(bytesData)
}
