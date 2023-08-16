package controller

import (
	"strconv"
	"web/service"
)

type InsertNewPostResponse struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 成功/错误信息
	Data interface{} `json:"data"` // 数据
}

func InsertNewPost(topicIdStr, content string) *InsertNewPostResponse {
	if len(content) > 300 {
		return &InsertNewPostResponse{
			Code: -1,
			Msg:  "content too long",
		}
	}

	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &InsertNewPostResponse{
			Code: -1,
			Msg:  "illegal topic id",
		}
	}

	postId, err := service.NewInsertPostFlow().Insert(topicId, content)
	if err != nil {
		return &InsertNewPostResponse{
			Code: -1,
			Msg:  "unknown error",
		}
	}

	return &InsertNewPostResponse{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{"postId": postId},
	}
}
