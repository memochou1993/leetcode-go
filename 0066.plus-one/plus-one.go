package problem0066

func plusOne(digits []int) []int {
	digits[len(digits)-1]++

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] < 10 {
			return digits
		}

		digits[i] = 0
		digits[i-1]++
	}

	if digits[0] == 10 {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}

	return digits
}
