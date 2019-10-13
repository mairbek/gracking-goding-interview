package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	// {
	// 	result := []int{1, 2, 3, 4, 5, 6, 7}
	// 	rotate(result, 3)
	// 	assert.Equal(t, result, []int{5, 6, 7, 1, 2, 3, 4})
	// }
	{
		result := []int{1, 2}
		rotate(result, 3)
		assert.Equal(t, result, []int{2, 1})
	}
}

func TestTwoSum(t *testing.T) {
	{

		assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	}
}

