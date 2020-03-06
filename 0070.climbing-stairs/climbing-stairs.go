package problem0070

func climbStairs(n int) int {
	a, b := 1, 1+n%2

	for i := 0; i < n/2; i++ {
		a += b
		b += a
	}

	return a
}
