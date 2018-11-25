package leetcode

import (
	"testing"
)

func fromNumber(in int) *ListNode {
	value := in % 10
	result := new(ListNode)
	result.Val = value
	current := result
	for in > 0 {
		in /= 10
		value := in % 10
		newlist := new(ListNode)
		newlist.Val = value

		current.Next = newlist
		current = newlist
	}
	return result
}

func toNumber(in *ListNode) int {
	result := 0
	multi := 1

	for in != nil {
		result += multi * in.Val
		in = in.Next
		multi *= 10
	}
	return result
}
func TestTimeConsuming(t *testing.T) {
	compare := func(a int, b int, expected int) {
		actual := toNumber(addTwoNumbers(fromNumber(a), fromNumber(b)))
		if actual != expected {
			t.Errorf("Expected %d got %d", expected, actual)
		}
	}
	compare(342, 465, 807)
	compare(0, 465, 465)
	compare(5, 6, 11)
	compare(0, 0, 0)
	compare(3061, 8036, 11097)

}
