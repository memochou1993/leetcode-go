# Reverse Integer

## Description

Given a 32-bit signed integer, reverse digits of an integer.

- Example 1:

```BASH
Input: 123
Output: 321
Example 2:
```

- Example 2:

```BASH
Input: -123
Output: -321
Example 3:
```

- Example 3:

```BASH
Input: 120
Output: 21
```

## Solution

```GO
func reverse(x int) int {
	// 宣告一個正負數的標記
	sign := 1

	// 如果 x 是負數，將標記設為 -1，並將 x 修改為正數
	if x < 0 {
		sign = -1
		x *= -1
	}

	result := 0
	for x > 0 {
		// 取得 x 的尾數
		temp := x % 10
		// 將 x 的尾數加至 result
		result = result*10 + temp
		// 去除 x 的尾數
		x = x / 10
	}

	// 用正負數的標記修正 result
	result *= sign

	// 避免 result 溢出
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result
}
```