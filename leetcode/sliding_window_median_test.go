package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowMedian(t *testing.T) {
	actual := medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)

	assert.InDeltaSlice(t, []float64{1, -1, -1, 3, 5, 6}, actual, 0.001)
}
