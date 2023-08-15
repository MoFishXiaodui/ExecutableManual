package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryPostById(t *testing.T) {
	_ = InitPostIndexMap("../data/")
	postDao := NewPostDaoInstance()
	res := postDao.QueryPostById(1)
	expect := &Post{Content: "小姐姐快来1"}
	assert.Equal(t, expect.Content, res.Content)
}

// TestPostDao_QueryPostsByParentId 通过比较获取的posts的切片长度来测试
func TestPostDao_QueryPostsByParentId(t *testing.T) {
	_ = InitPostIndexMap("../data/")
	postDao := NewPostDaoInstance()
	res := postDao.QueryPostsByParentId(1)
	expect := 5
	log.Println(res[2].Content)
	assert.Equal(t, expect, len(res))
}
