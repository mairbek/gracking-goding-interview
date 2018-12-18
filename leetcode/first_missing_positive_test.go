package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
  {

	  assert.Equal(t, 3, firstMissingPositive([]int{1,2,0}))
	  assert.Equal(t, 2, firstMissingPositive([]int{3,4,-1,1}))
	  assert.Equal(t, 1, firstMissingPositive([]int{7,8,9,11,12}))
	  assert.Equal(t, 2, firstMissingPositive([]int{1, 1}))
  }
}
