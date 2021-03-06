# 136. Single Number

## Description

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

- Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

- Example 1:

```BASH
Input: [2,2,1]
Output: 1
```

- Example 2:

```BASH
Input: [4,1,2,1,2]
Output: 4
```

## Solution

```GO
func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		// 使用 XOR 運算符找出唯一的位元
		res ^= num
	}

	return res
}
```

## Note

假設有以下參數：

```BASH
nums: [4, 1, 2, 1, 2]
```

說明：

```BASH
取得 0 和某一個位元的 XOR 時，它會回傳該位元：

a⊕0=a

取得兩個一樣的位元的 XOR 時，它會回傳 0：

a⊕a=0

因此，可以使用 XOR 找出所有位元中唯一的數字：

a⊕b⊕a=(a⊕a)⊕b=0⊕b=b

此陣列可以看成：

4⊕(1⊕1)⊕(2⊕2)=4⊕0⊕0=4

最終返回：4
```
