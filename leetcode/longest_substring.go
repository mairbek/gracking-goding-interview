package leetcode

func lengthOfLongestSubstring(s string) int {
	visited := make(map[byte]int)
	longest := 0
	start := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if pos, ok := visited[c]; ok {
			current := i - start
			if current > longest {
				longest = current
			}
			if pos >= start {
				start = pos + 1
			}
		}
		visited[c] = i
	}
	current := len(s) - start
	if current > longest {
		longest = current
	}
	return longest
}

func bruteForce(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		current := 0
		visited := make(map[byte]int)
		for j := i; j < len(s); j++ {
			c := s[j]
			if _, ok := visited[c]; ok {
				if current > longest {
					longest = current
					current = 0
				}
				break
			}
			current++
			visited[c] = i
		}
		if current > longest {
			longest = current
		}
	}
	return longest
}
