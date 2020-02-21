package problem0013

import "log"

func romanToInt(s string) int {
	result := 0

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	log.Println(m)

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]

		sign := 1
		if temp < last {
			sign = -1
		}

		result += sign * temp

		last = temp
	}

	return result
}
