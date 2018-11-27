package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.InDelta(t, mf.FindMedian(), 1.5, 0.001)

	mf.AddNum(3)
	assert.InDelta(t, mf.FindMedian(), 2., 0.001)

	mf = Constructor()
	mf.AddNum(-1)
	mf.AddNum(-2)
	mf.AddNum(-3)
	mf.AddNum(-4)
	mf.AddNum(-5)

	assert.InDelta(t, mf.FindMedian(), -3., 0.001)
}
