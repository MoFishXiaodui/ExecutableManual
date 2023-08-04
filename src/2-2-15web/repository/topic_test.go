package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitTopicIndexMap(t *testing.T) {
	var expect error = nil
	output := InitTopicIndexMap("../data/")
	assert.Equal(t, expect, output)
}

func TestQueryTopicById(t *testing.T) {
	_ = InitTopicIndexMap("../data/")
	topicDao := NewTopicDaoInstance()
	res := topicDao.QueryTopicById(1)
	expect := &Topic{Content: "冲冲冲！"}
	assert.Equal(t, expect.Content, res.Content)
}
