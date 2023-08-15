package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

// PostDao 对外提供PostDao模型
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

// 快速通过id找到相应的post的map表
var postIndexMap map[int64]*Post

// 快速通过parentId找到相应的posts列表的map表
var postParentIdMap map[int64][]*Post

// InitPostIndexMap 同时初始化postIndexMap和postParentIdMap
func InitPostIndexMap(filePath string) error {
	file, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	postTmpMap := make(map[int64]*Post)
	postParentTmpMap := make(map[int64][]*Post)

	for scanner.Scan() {
		text := scanner.Text()
		post := Post{}
		err = json.Unmarshal([]byte(text), &post)
		if err != nil {
			return err
		}

		postTmpMap[post.Id] = &post

		posts, ok := postParentTmpMap[post.ParentId]
		if ok {
			postParentTmpMap[post.ParentId] = append(posts, &post)
		} else {
			postParentTmpMap[post.ParentId] = []*Post{&post}
		}
	}

	postIndexMap = postTmpMap
	postParentIdMap = postParentTmpMap

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() {
		postDao = &PostDao{}
	})
	return postDao
}

func (*PostDao) QueryPostById(id int64) *Post {
	return postIndexMap[id]
}

func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	return postParentIdMap[parentId]
}
