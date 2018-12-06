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

	matchers := make([]Matcher, 0)
	mlen := 0
	for _, m := range p {
		if m == rune('.') {
			matchers = append(matchers, Matcher{any: true, all: false})
			mlen++
		} else if m == rune('*') {
			matchers[mlen-1].all = true
		} else {
			matchers = append(matchers, Matcher{char: m, any: false, all: false})
			mlen++
		}
	}

	return runMatch(matchers, s, 0, 0)
}

func runMatch(matchers []Matcher, s []rune, mpos int, spos int) bool {
	if mpos >= len(matchers) && spos >= len(s) {
		return true
	}
	if mpos >= len(matchers) {
		return false
	}
	m := matchers[mpos]
	isMatch := spos < len(s) && (m.any || m.char == s[spos])
	if isMatch {
		if runMatch(matchers, s, mpos+1, spos+1) {
			return true
		}
		if m.all {
			if runMatch(matchers, s, mpos, spos+1) {
				return true
			}
		}
	}
	if m.all {
		if runMatch(matchers, s, mpos+1, spos) {
			return true
		}
	}
	return false
}
