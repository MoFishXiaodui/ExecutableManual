package repository

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	_ = InitPostIndexMap("../data/")
	_ = InitTopicIndexMap("../data/")
	os.Exit(m.Run())
}

func TestQueryPostById(t *testing.T) {
	postDao := NewPostDaoInstance()
	res := postDao.QueryPostById(1)
	expect := &Post{Content: "小姐姐快来1"}
	assert.Equal(t, expect.Content, res.Content)
}

// TestPostDao_QueryPostsByParentId 通过比较获取的posts的切片长度来测试
func TestPostDao_QueryPostsByParentId(t *testing.T) {
	postDao := NewPostDaoInstance()
	res := postDao.QueryPostsByParentId(1)
	expect := 5
	log.Println(res[2].Content)
	assert.Equal(t, expect, len(res))
}

func TestPostDao_InsertNewPost(t *testing.T) {
	postDao := NewPostDaoInstance()
	err := postDao.insertNewPost(12, 4, "你好呀可爱的小伙伴", time.Now().Unix())
	if err != nil {
		log.Println("err", err)
	}
	post := postDao.QueryPostById(12)
	posts := postDao.QueryPostsByParentId(4)
	fmt.Printf("post->%#v\ntopic:->%#v\n", post, posts)
}

func TestPostDao_NewPost(t *testing.T) {
	postDao := NewPostDaoInstance()
	_, err := postDao.NewPost(2, "测试自动补充postid和时间戳")
	if err != nil {
		_ = fmt.Errorf("err: %v", err)
	}
}
