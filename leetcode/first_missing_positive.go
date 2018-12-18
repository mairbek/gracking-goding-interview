package leetcode

// Given an unsorted integer array, find the smallest missing positive integer.
// Requires O(n) time, O(1) space.

func firstMissingPositive(nums []int) int {
  // The idea is to reuse input array to store all the required information.

  // First, we partition the array and put all the positive integers in the begining.
  pos := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] > 0 {
      tmp := nums[pos]
      nums[pos] = nums[i]
      nums[i] = tmp
      pos++
    }
  }
  // Now we mark all the positive elements in the array: if nums[i] < 0, then i is present in the array.
  for i := 0; i < pos; i++ {
    k := nums[i]
    if k < 0 {
      k *= -1
    }
    // Too large
    if k > pos {
      continue
    }
    if nums[k-1] > 0 {
      nums[k-1] *= -1
    }
  }
  // Now we find a first positive integer.
  for i := 0; i < pos; i++ {
    if nums[i] > 0 {
      return i + 1
    }
  }
  return pos + 1
}
