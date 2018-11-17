package main

import "fmt"

// Given a set of possible change, such as {1, 5, 10, 25}, and an amount such as 13, return all possible ways you can give change back to someone
func Change(num int, coins []int) int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		count := 0
		for _, coin := range coins {
			if coin == i {
				count++
				fmt.Printf("%d %d %d \n", coin, i, count)
			} else if coin < i {
				count += result[i-coin]
				fmt.Printf("%d %d %d \n", coin, i, count)
			}
		}
		result[i] = count
	}
	fmt.Printf("%v", result)
	return result[num]
}
