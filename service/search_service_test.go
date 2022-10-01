package service

import (
	"fmt"
	"testing"
)

func TestSearchService_Search(t *testing.T) {
	service := &SearchService{
		S:    "测试",
		Type: "1",
	}
	code, resp := service.Search()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
