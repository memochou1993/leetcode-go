package problem0125

func isPalindrome(s string) bool {
	arr := []rune{}

	for _, b := range s {
		if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			arr = append(arr, b)
		}
	}

	for i := 0; i < len(arr)/2; i++ {
		diff := arr[i] - arr[len(arr)-i-1]

		if diff == 0 {
			continue
		}

		if arr[i] > '9' && diff != 32 && diff != -32 {
			return false
		}

		if arr[i] <= '9' && diff != 0 {
			return false
		}
	}

	return true
}
