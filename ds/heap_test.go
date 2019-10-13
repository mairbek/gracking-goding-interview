package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := NewMinHeap()

	assertMin := func(expected int) {
		ok, min := h.Min()
		assert.True(t, ok)
		assert.Equal(t, expected, min)
	}

	assertExtractMin := func(expected int) {
		ok, min := h.ExtractMin()
		assert.True(t, ok)
		assert.Equal(t, expected, min)
	}

	assertEmpty := func() {
		ok, _ := h.Min()
		assert.False(t, ok)
	}

	assertEmpty()

	h.Insert(2)
	assertMin(2)

	h.Insert(1)
	assertMin(1)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	h.Insert(8)
	assertMin(1)

	h.Insert(0)
	assertMin(0)

	assertExtractMin(0)
	assertExtractMin(1)
	assertExtractMin(2)
	assertExtractMin(5)

	h.Insert(0)
	assertMin(0)
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	assertMin := func(expected int) {
		ok, min := h.Min()
		assert.True(t, ok)
		assert.Equal(t, expected, min)
	}

	assertExtractMin := func(expected int) {
		ok, min := h.ExtractMin()
		assert.True(t, ok)
		assert.Equal(t, expected, min)
	}

	assertEmpty := func() {
		ok, _ := h.Min()
		assert.False(t, ok)
	}

	assertEmpty()

	h.Insert(2)
	assertMin(2)

	h.Insert(1)
	assertMin(2)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	h.Insert(8)
	assertMin(8)

	assertExtractMin(8)
	assertExtractMin(7)
	assertExtractMin(6)
	assertExtractMin(5)
}
