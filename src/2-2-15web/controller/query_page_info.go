package controller

import (
	"log"
	"strconv"
	"web/service"
)

type PageData struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 成功/错误信息
	Data interface{} `json:"data"` // 数据
}

func QueryPageInfo(topicIdStr string) *PageData {
	// 将string类型的topicIdStr转换为int64类型的topicId
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	// ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
	// The string may begin with a leading sign: "+" or "-".
	// If the base argument is 0, the true base is implied by the string's prefix following the sign (if present):
	// 2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise.
	// Also, for argument base 0 only, underscore characters are permitted as defined by the Go syntax for integer literals.
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	pageInfo, err := service.QueryPageInfo(topicId)
	log.Println("getPageInfo", topicId, pageInfo.PostList)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}
}
