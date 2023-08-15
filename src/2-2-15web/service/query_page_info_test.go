package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"web/repository"
)

func TestMain(m *testing.M) {
	_ = repository.InitTopicIndexMap("../data/")
	_ = repository.InitPostIndexMap("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.Equal(
		t,
		PageInfo{Topic: &repository.Topic{Content: "冲冲冲！"}}.Topic.Content,
		pageInfo.Topic.Content,
	)
}

func TestQueryPageInfoForPost(t *testing.T) {
	pageInfo, _ := QueryPageInfo(2)
	assert.Equal(
		t,
		5,
		len(pageInfo.PostList),
	)
}
