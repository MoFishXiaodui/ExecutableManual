package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
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
	postDao      *PostDao
	postFilePath string
	postOnce     sync.Once
)

// 快速通过id找到相应的post的map表
var postIndexMap map[int64]*Post

// 快速通过parentId找到相应的posts列表的map表
var postParentIdMap map[int64][]*Post

// InitPostIndexMap 同时初始化postIndexMap和postParentIdMap
func InitPostIndexMap(filePath string) error {
	postFilePath = filePath + "post"
	file, err := os.Open(postFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	postTmpMap := make(map[int64]*Post)
	postParentTmpMap := make(map[int64][]*Post)

	// 记录占用id的情况（只需要记录最大值即可）
	tmpMaxPostId := int64(1)

	for scanner.Scan() {
		text := scanner.Text()
		post := Post{}
		err = json.Unmarshal([]byte(text), &post)
		if err != nil {
			return err
		}

		if post.Id > tmpMaxPostId {
			tmpMaxPostId = post.Id
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
	newIdDistributeInstance().updatePostId(tmpMaxPostId)

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

func (*PostDao) insertNewPostCheckParam(id, parentId int64) bool {
	// 先判断是不是用新的id来标识post
	if _, ok := postIndexMap[id]; ok {
		return false
	}
	// 在判断是不是已有topic可以关联
	if _, ok := topicIndexMap[parentId]; ok {
		return true
	}
	return false
}

// insertNewPost - 插入新的Post，需要自己设定postId和createTime时间戳。一般情况下应该使用 NewPost 函数，而非本函数。
func (dao *PostDao) insertNewPost(id, parentId int64, content string, createTime int64) error {
	if !dao.insertNewPostCheckParam(id, parentId) {
		return errors.New("插入post时参数校验不通过")
	}
	file, err := os.OpenFile(postFilePath, os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Println("无法打开post数据文件:", err)
		return err
	}
	defer file.Close()

	newPost := Post{
		Id:         id,
		ParentId:   parentId,
		Content:    content,
		CreateTime: createTime,
	}
	newTextLine, err := json.Marshal(newPost)
	if err != nil {
		log.Println("post数据转化为json出错")
		return err
	}

	// 写入文件
	log.Println(string(newTextLine))
	_, err = file.WriteString("\n" + string(newTextLine))
	if err != nil {
		fmt.Println("无法写入post文件:", err)
		return err
	}

	// 刷新postIndexMap
	postIndexMap[newPost.Id] = &newPost

	// 刷新postParentIdMap
	posts, ok := postParentIdMap[newPost.ParentId]
	if ok {
		postParentIdMap[newPost.ParentId] = append(posts, &newPost)
	} else {
		postParentIdMap[newPost.ParentId] = []*Post{&newPost}
	}

	// 刷新postId占用情况
	newIdDistributeInstance().updatePostId(newPost.Id)

	return nil
}

// NewPost - 生成新的Post并插入到文件中。自动生成postId和时间戳
func (dao *PostDao) NewPost(parentId int64, content string) (int64, error) {
	postId := newIdDistributeInstance().generatePostId()
	createTime := time.Now().Unix()
	return postId, dao.insertNewPost(postId, parentId, content, createTime)
}
