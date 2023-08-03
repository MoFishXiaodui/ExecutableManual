package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}
