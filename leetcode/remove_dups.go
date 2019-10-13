package leetcode

func removeDuplicates(nums []int) int {
	removed := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			removed++
		}
	}
	return 0
}
