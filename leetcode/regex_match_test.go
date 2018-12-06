package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	assert.True(t, isMatch("a", "a"))
	assert.False(t, isMatch("a", "b"))
	assert.True(t, isMatch("a", "."))
	assert.False(t, isMatch("bab", "aba"))
	assert.True(t, isMatch("aba", "aba"))
	assert.True(t, isMatch("aba", "a.a"))
	assert.False(t, isMatch("aba", "a"))
	assert.True(t, isMatch("aaaa", "a*"))
	assert.True(t, isMatch("aab", "c*a*b"))
	assert.False(t, isMatch("ab", ".*c"))
	assert.True(t, isMatch("", "c*"))
	assert.False(t, isMatch("a", ""))
	assert.True(t, isMatch("a", "ab*"))
	assert.True(t, isMatch("bbbba", ".*a*a"))
}
