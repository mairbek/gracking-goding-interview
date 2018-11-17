package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	root := AsList([]int{1, 2, 3, 4, 5})
	PrintList(root)

	fmt.Printf("Middle element %d\n", MiddleNode(root))
	PrintList(ReverseList(root))
}
