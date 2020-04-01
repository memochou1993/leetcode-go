package problem0101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	pre []int
	in  []int
}

type answer struct {
	first bool
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				pre: []int{1, 2, 2},
				in:  []int{2, 1, 2},
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				pre: []int{1, 2, 3, 2, 3},
				in:  []int{3, 2, 1, 2, 3},
			},
			a: answer{
				first: true,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, isSymmetric(ltot(p.pre, p.in)))
	}
}

// convert []int to *TreeNode
func ltot(pre []int, in []int) *TreeNode {
	if len(pre) != len(in) {
		return nil
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	index := indexOf(res.Val, in)

	res.Left = ltot(pre[1:index+1], in[:index])
	res.Right = ltot(pre[index+1:], in[index+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return -1
}
