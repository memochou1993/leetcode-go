package problem0108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first []int
}

type answer struct {
	first []int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: []int{1, 2, 3, 4, 5, 6, 7},
			},
			a: answer{
				first: []int{4, 2, 1, 3, 6, 5, 7},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, ptol(sortedArrayToBST(p.first)))
	}
}

// convert preorder *TreeNode to []int
func ptol(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, ptol(root.Left)...)
	res = append(res, ptol(root.Right)...)

	return res
}
