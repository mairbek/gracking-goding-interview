package leetcode

// ListNode is a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AsList(vals []int) *ListNode {
	var root *ListNode
	var current *ListNode
	for _, val := range vals {
		if root == nil {
			root = new(ListNode)
			current = root
		} else {
			current.Next = new(ListNode)
			current = current.Next
		}
		current.Val = val
	}
	return root
}

func AsArray(node *ListNode) []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
