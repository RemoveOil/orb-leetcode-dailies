package kthlargestlevelsum

import (
	"solutions/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	assert.Equal(
		t,
		int64(13),
		kthLargestLevelSum(
			objects.ParseTreeFromStringMust("[5,8,9,2,1,3,7,4,6]"),
			2,
		),
	)
	assert.Equal(
		t,
		int64(3),
		kthLargestLevelSum(
			objects.ParseTreeFromStringMust("[1,2,null,3]"),
			1,
		),
	)
	assert.Equal(
		t,
		int64(-1),
		kthLargestLevelSum(
			objects.ParseTreeFromStringMust("[5,8,9,2,1,3,7]"),
			4,
		),
	)
}
