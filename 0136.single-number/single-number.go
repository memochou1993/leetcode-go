package problem0136

func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res ^= num
	}

	return res
}
