package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	result := make([]int, n)

	for i := 0; i < n; i++ {
		j := (i - k + n) % n
		result[i] = nums[j]
	}

	for i := 0; i < n; i++ {
		nums[i] = result[i]
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
