package problem0067

func addBinary(a string, b string) string {
	la := len(a) - 1
	lb := len(b) - 1

	temp := 0
	carry := 0
	result := ""

	for la >= 0 || lb >= 0 || carry != 0 {
		temp = carry

		if la >= 0 {
			temp += int(a[la] - byte('0'))
			la--
		}

		if lb >= 0 {
			temp += int(b[lb] - byte('0'))
			lb--
		}

		result = string(temp%2+'0') + result

		carry = temp / 2
	}

	return result
}
