package controller

import "Go_server/define"

func NewQueryRequest() *define.QueryRequest {
	return &define.QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		Keyword: "",
		Status:  -1,
	}
}
