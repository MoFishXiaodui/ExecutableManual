package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// 1-先声明一个结构体，json导出和 data/topic 的内容 对应上
type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

// 4-对外提供Dao模型
type TopicDao struct {
}

// 5-设置锁
var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

// 2-创建索引表。可以快速通过Id值找到对应的Topic
var topicIndexMap map[int64]*Topic

// 3-InitTopicIndexMap 初始化话题索引
// 如果要规范的话，后面要设置成不可导出（即小写字母开头，但是为了方便测试，先用着大写字母开头）
func InitTopicIndexMap(filePath string) error {
	// 尝试打开文件
	file, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	// 为临时topic索引分配空间
	topicTmpMap := make(map[int64]*Topic)

	fmt.Println("开始读取数据")

	// 读取文本
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)

		// 新建单个Topic结构
		var topic Topic

		// 转换text填充数据到topic上
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}

		// 把topic指针添加到临时索引表上
		topicTmpMap[topic.Id] = &topic
	}
	// 把 临时索引表 赋给 最终的索引表
	topicIndexMap = topicTmpMap

	return nil
}

// 6-定义 TopicDao 初始函数
func NewTopicDaoInstance() *TopicDao {
	// func (*sync.Once).Do(f func())
	// Once is an object that will perform exactly one action.
	// A Once must not be copied after first use.
	topicOnce.Do(
		func() {
			// 创建TopicDao实例并把该实例的指针赋值给全局的topicDao
			topicDao = &TopicDao{}
		})
	return topicDao
}

// 7-定义topic查询函数
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
