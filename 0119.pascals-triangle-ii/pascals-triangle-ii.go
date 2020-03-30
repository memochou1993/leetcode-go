package problem0119

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)

	res[0] = 1
	res[rowIndex] = 1

	mid := (rowIndex+1)/2 + 1

	for i := 1; i < mid; i++ {
		res[i] = res[i-1] * (rowIndex + 1 - i) / i
		res[rowIndex-i] = res[i]
	}

	return res
}
