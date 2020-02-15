package main

import (
	"fmt"
)

func main() {
	result := twoSum([]int{1, 2, 3}, 5)

	fmt.Println(result)
}

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
