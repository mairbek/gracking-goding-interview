package leetcode

type Matcher struct {
	char rune
	any  bool
	all  bool
}

type ToMatch struct {
	matcherPos int
	stringPos  int
	next       *ToMatch
}

func isMatch(ss string, ps string) bool {
	p := []rune(ps)
	s := []rune(ss)

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == rune('*') {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == rune('.')) && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == rune('.')) && dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
