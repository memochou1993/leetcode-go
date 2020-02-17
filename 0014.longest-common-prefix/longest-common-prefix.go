package problem0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][:i+1]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][:i+1] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
