package service

import (
	"fmt"
	"testing"
)

func TestSearchService_Search(t *testing.T) {
	service := &SearchService{
		S:    "周杰伦",
		Type: "100",
	}
	code, resp := service.Search()
	fmt.Println(code, string(resp))
	if code != 200 {
		t.Errorf("code error: %f", code)
	}
}
