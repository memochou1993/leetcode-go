package problem0058

func lengthOfLastWord(s string) int {
	cursor := len(s)

	for ; cursor > 0; cursor-- {
		if s[cursor-1] != byte(' ') {
			break
		}
	}

	s = s[:cursor]

	for ; cursor > 0; cursor-- {
		if s[cursor-1] == byte(' ') {
			break
		}
	}

	s = s[cursor:]

	return len(s)
}
