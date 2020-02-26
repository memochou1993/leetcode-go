package problem0053

func maxSubArray(nums []int) int {
	max := -(1 << 63)
	temp := 0

	for _, v := range nums {
		if temp+v < v {
			temp = v
		} else {
			temp += v
		}

		if temp > max {
			max = temp
		}
	}

	return max
}
