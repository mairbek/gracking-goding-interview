package leetcode

type ListHeap struct {
	Values []*ListNode
	Len    int
}

func (h *ListHeap) Min() (bool, *ListNode) {
	if h.Len > 0 {
		return true, h.Values[0]
	}
	return false, nil
}

func (h *ListHeap) ExtractMin() (bool, *ListNode) {
	if h.Len > 0 {
		result := h.Values[0]
		h.Values[0] = h.Values[h.Len-1]
		h.Values[h.Len-1] = nil
		h.Len--
		h.topDown()
		return true, result
	}
	return false, nil
}

func (h *ListHeap) Insert(value *ListNode) {
	if h.Len < len(h.Values) {
		h.Values[h.Len] = value
		h.Len++
	} else {
		h.Values = append(h.Values, value)
		h.Len++
	}
	h.bottomUpRebalance()
}

func (h *ListHeap) bottomUpRebalance() {
	pos := h.Len - 1
	for pos > 0 {
		parent := (pos - 1) / 2
		if h.Values[parent].Val <= h.Values[pos].Val {
			break
		}
		h.Values[parent], h.Values[pos] = h.Values[pos], h.Values[parent]
		pos = parent
	}
}

func (h *ListHeap) topDown() {
	pos := 0
	for {
		val := h.Values[pos]
		swap := -1
		one := 2*pos + 1
		if one < h.Len && val.Val > h.Values[one].Val {
			val = h.Values[one]
			swap = one
		}
		two := 2*pos + 2
		if two < h.Len && val.Val > h.Values[two].Val {
			swap = two
		}
		if swap < 0 {
			break
		}
		h.Values[pos], h.Values[swap] = h.Values[swap], h.Values[pos]
		pos = swap
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := new(ListHeap)
	var result *ListNode
	var current *ListNode

	for _, list := range lists {
		// Yes, fuc u leetcode
		if list == nil {
			continue
		}
		heap.Insert(list)
	}

	for ok, item := heap.ExtractMin(); ok; ok, item = heap.ExtractMin() {
		if current == nil {
			current = item
			result = item
		} else {
			current.Next = item
			current = item
		}
		if item.Next != nil {
			heap.Insert(item.Next)
		}
	}
	return result
}
