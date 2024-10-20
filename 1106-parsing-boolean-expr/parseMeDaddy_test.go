package parsingbooleanexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, false, parseBoolExpr("|(f,f,f,f)"))
	assert.Equal(t, true, parseBoolExpr("|(f,t,f,f)"))
	assert.Equal(t, false, parseBoolExpr("f"))
	assert.Equal(t, true, parseBoolExpr("t"))
	assert.Equal(t, false, parseBoolExpr("&(|(f))"))
	assert.Equal(t, true, parseBoolExpr("!(&(f,t))"))
	assert.Equal(t, false, parseBoolExpr("!(|(f,t))"))
}
