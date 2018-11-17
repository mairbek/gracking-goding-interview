package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

// AsList converts array to list, usualy to test.
func AsList(vals []int) *Node {
	var root *Node
	var current *Node
	for _, val := range vals {
		if root == nil {
			root = new(Node)
			current = root
		} else {
			current.next = new(Node)
			current = current.next
		}
		current.value = val
	}
	return root
}

// PrintList prints list, usually to test.
func PrintList(root *Node) {
	current := root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

// MiddleNode returns the middle node of the list.
func MiddleNode(root *Node) int {
	current := root
	middle := root
	count := 0
	for current != nil {
		if count > 0 && count%2 == 0 {
			middle = middle.next
		}
		current = current.next
		count++
	}
	return middle.value
}

func ReverseList(root *Node) *Node {
	var result *Node
	current := root
	for ; current != nil; current = current.next {
		previous := result
		result = new(Node)
		result.next = previous
		result.value = current.value
	}
	return result
}
