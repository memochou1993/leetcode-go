package problem0007

import (
	"math"
)

func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x *= -1
	}

	result := 0
	for x > 0 {
		temp := x % 10
		result = result*10 + temp
		x = x / 10
	}

	result *= sign

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result
}
