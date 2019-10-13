package leetcode

import (
	"sort"
	"strings"
)

// an O(n) alternative to sort is to count number of elements in the array.
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range(strs) {
		key := SortString(s)
		vals, ok := m[key]
		if ok {
			m[key] = append(vals, s)
		} else {
			m[key] = []string{s}
		}
	}
	result := make([][]string, len(m))
	i := 0
	for _, vals := range(m) {
		result[i] = vals
		i++
	}
	return result
}
