package problem0038

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)

	result := []byte{}
	count := 0
	last := byte('0')

	for i := 0; i < len(s); i++ {
		if last == byte('0') {
			last = s[i]
			count++
			continue
		}

		if last == s[i] {
			count++
			continue
		}

		result = append(result, byte(count+'0'), last)

		last = s[i]
		count = 1
	}

	result = append(result, byte(count+'0'), last)

	return string(result)
}
