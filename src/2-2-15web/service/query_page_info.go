package service

import (
	"errors"
	"sync"
	"web/repository"
)

// PageInfo 是最终要返回给上层(controller)函数的实体
type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

// QueryPageInfoFlow 处理请求内容和生成数据的结构体
type QueryPageInfoFlow struct {
	// 接收上层传来的topicId
	topicId int64

	// 组装的PageInfo实体，用来返回给上层
	pageInfo *PageInfo

	// 下面是获取散装的模型数据
	topic *repository.Topic // topic

	posts []*repository.Post
}

// 用来检查参数的函数
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

// 用来通过底层模型获取对应数据的函数
func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)

	// 获取topic信息
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()

	// 获取post信息
	go func() {
		defer wg.Done()
		f.posts = repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
	}()

	wg.Wait()
	return nil
}

// 用来组装成PageInfo实体的函数
func (f *QueryPageInfoFlow) preparePageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}

// QueryPageInfo 外部来了一个请求，我们创建一个新的QueryPageInfoFlow结构来处理，这个结构的Do()方法最终返回想要的结果
func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

// NewQueryPageInfoFlow 创建新QueryPageInfoFlow实例的函数
func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

// Do Do函数封装了检查参数、获取数据、组装实体三个步骤，并返回PageInfo
func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.preparePageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}
