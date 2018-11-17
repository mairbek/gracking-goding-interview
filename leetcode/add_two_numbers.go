package leetcode

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

// ListNode is a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	leftover := 0

	result := new(ListNode)
	result.Val = 0
	var current *ListNode

	for {
		if l1 != nil {
			leftover += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			leftover += l2.Val
			l2 = l2.Next
		}
		if leftover == 0 && l1 == nil && l2 == nil {
			break
		}
		if current == nil {
			current = new(ListNode)
			result = current
		} else {
			newNode := new(ListNode)
			current.Next = newNode
			current = newNode
		}
		current.Val = leftover % 10
		leftover /= 10
	}
	return result
}
