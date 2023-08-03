package coverage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudge(t *testing.T) {
	// score1 := Judge(66)
	// assert.Equal(t, true, score1)

	score2 := Judge(40)
	assert.Equal(t, false, score2)
}
