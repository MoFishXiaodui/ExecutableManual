package main

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line11"
	})
	defer monkey.Unpatch(ReadFirstLine)

	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}
