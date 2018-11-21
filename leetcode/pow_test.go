package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	compare := func(a float64, b int, expected float64) {
		actual := myPow(a, b)
		if actual != expected {
			assert.InDelta(t, actual, expected, .001)
		}
	}
	compare(2, 0, 1)
	compare(2, 10, 1024)
	compare(2.1, 3, 9.261000)
	compare(2., -2, 0.25)
	compare(2., -5, 0.03125)
}
