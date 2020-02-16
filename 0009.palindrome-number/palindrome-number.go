package problem0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	copy := x
	reverse := 0

	for copy > 0 {
		reverse = reverse*10 + copy%10
		copy /= 10
	}

	return x == reverse
}
