package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	compare := func(values [][]int, expected []int) {
		lists := make([]*ListNode, len(values))
		for i := 0; i < len(values); i++ {
			lists[i] = AsList(values[i])
		}
		result := mergeKLists(lists)
		actual := AsArray(result)
		assert.ElementsMatch(t, actual, expected)
	}
	compare([][]int{
		[]int{1, 4, 5},
		[]int{1, 3, 4},
		[]int{2, 6},
	}, []int{1, 1, 2, 3, 4, 4, 5, 6})
}
