package leetcode

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	compare := func(s string, expected int) {
		actual := lengthOfLongestSubstring(s)
		if actual != expected {
			t.Errorf("Expected %d got %d", expected, actual)
		}
	}

	compare("", 0)
	compare(" ", 1)
	compare("abcabcbb", 3)
	compare("bbbb", 1)
	compare("pwwkew", 3)
	compare("abba", 2)
	compare("dvdf", 3)
}
