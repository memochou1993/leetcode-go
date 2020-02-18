package problem0020

func isValid(s string) bool {
	stack := []string{}

	pairs := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for i := 0; i < len(s); i++ {
		char := s[i : i+1]

		if opposite, ok := pairs[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opposite {
				return false
			}

			stack = stack[:len(stack)-1]

			continue
		}

		stack = append(stack, char)
	}

	return len(stack) == 0
}
