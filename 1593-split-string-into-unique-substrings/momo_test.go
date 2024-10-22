package splitstringintouniquesubstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	assert.Equal(t, 5, maxUniqueSplit("ababccc"))
	assert.Equal(t, 1, maxUniqueSplit("aa"))
	assert.Equal(t, 2, maxUniqueSplit("aaaa"))
	assert.Equal(t, 6, maxUniqueSplit("abcdef"))
}
