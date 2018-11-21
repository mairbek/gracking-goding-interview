package leetcode

func myPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		neg = true
		n *= -1
	}
	result := 1.
	for k := n; k > 0; {
		if k%2 == 1 {
			if neg {
				result /= x
			} else {
				result *= x
			}
			k--
		} else {
			x *= x
			k /= 2
		}
	}
	return result
}
