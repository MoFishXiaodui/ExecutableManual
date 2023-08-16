package service

import "web/repository"

type InsertPostFlow struct {
}

func NewInsertPostFlow() *InsertPostFlow {
	return &InsertPostFlow{}
}

// Insert 插入新的Post，需传入回帖内容并关联话题id。成功时返回对应的postId和nil值。
func (f *InsertPostFlow) Insert(topicId int64, content string) (postId int64, err error) {
	postId, err = repository.NewPostDaoInstance().NewPost(topicId, content)
	return postId, err
}
