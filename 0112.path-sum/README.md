# 112. Path Sum

## Description

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

- Note:

A leaf is a node with no children.

- Example:

Given the below binary tree and sum = 22,

```BASH
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
```

return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

## Solution

```GO
func hasPathSum(root *TreeNode, sum int) bool {
	// 如果根為空，則返回 false
	if root == nil {
		return false
	}

	// 將總數減掉節點的值
	sum -= root.Val

	// 判斷是否為葉
	if root.Left == nil && root.Right == nil {
		// 判斷總數是否被消去
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```

## Note

假設有以下參數：

```BASH
root: [5, 4, 11, 7, 2, 8, 13, 4, 1]
sum: 22
```

說明：

```BASH
樹的形狀如下：

--------------------
         5
        / \
       4   8
      /   / \
     11  13  4
    /  \      \
   7    2      1
--------------------

走遍每一個路徑，當路徑為 5->4->11->2 時，總和為 22。

最終返回：true
```
