package problem0001

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := index[target-num]; ok == true {
			return []int{j, i}
		}
		index[num] = i
	}

	return []int{}
}
