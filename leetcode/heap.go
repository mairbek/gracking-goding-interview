package leetcode

// Heap vs Binary tree:
// - Heap is implemented as a array so no memory overhead. For each node in the binary tree two pointers are allocated.
// - **Average** insertion in heap is O(1) comparing to O(logn).
// - Heap only supports min operation and is simpler than binary tree.
// - Heap is easier to implement, balancing red-black trees is annoying.
// - Creating heap is O(N) while creating binary tree is O(n*logn)
// Detailed summary here https://stackoverflow.com/questions/6147242/heap-vs-binary-search-tree-bst
type Heap struct {
	Values []int
	Len    int
}

func (h *Heap) Min() (bool, int) {
	if h.Len > 0 {
		return true, h.Values[0]
	}
	return false, -1
}

func (h *Heap) ExtractMin() (bool, int) {
	if h.Len > 0 {
		result := h.Values[0]
		h.Values[0] = h.Values[h.Len-1]
		h.Values[h.Len-1] = -1
		h.Len--
		h.topDown()
		return true, result
	}
	return false, -1
}

func (h *Heap) Insert(value int) {
	if h.Len < len(h.Values) {
		h.Values[h.Len] = value
		h.Len++
	} else {
		h.Values = append(h.Values, value)
		h.Len++
	}
	h.bottomUpRebalance()
}

func (h *Heap) bottomUpRebalance() {
	pos := h.Len - 1
	for pos > 0 {
		parent := (pos - 1) / 2
		if h.Values[parent] <= h.Values[pos] {
			break
		}
		h.Values[parent], h.Values[pos] = h.Values[pos], h.Values[parent]
		pos = parent
	}
}

func (h *Heap) topDown() {
	pos := 0
	for {
		val := h.Values[pos]
		swap := -1
		one := 2*pos + 1
		if one < h.Len && val > h.Values[one] {
			val = h.Values[one]
			swap = one
		}
		two := 2*pos + 2
		if two < h.Len && val > h.Values[two] {
			swap = two
		}
		if swap < 0 {
			break
		}
		h.Values[pos], h.Values[swap] = h.Values[swap], h.Values[pos]
		pos = swap
	}
}
