package controller

import (
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
	"web/repository"
)

func TestMain(m *testing.M) {
	repository.InitTopicIndexMap("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	pageData := QueryPageInfo("1")
	assert.Equal(
		t,
		int64(0),
		pageData.Code,
	)
}
