# Longest Common Prefix

## Description

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

- Example 1:

```BASH
Input: ["flower","flow","flight"]
Output: "fl"
```

- Example 2:

```BASH
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

## Solution

```GO
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		// 用第一個字串當作基準，從第一個字母開始比對
		char := strs[0][:i+1]

		// 從第二個字串開始比對
		for j := 1; j < len(strs); j++ {
			// 判斷比對的字母是否不一樣
			if i == len(strs[j]) || strs[j][:i+1] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
```
