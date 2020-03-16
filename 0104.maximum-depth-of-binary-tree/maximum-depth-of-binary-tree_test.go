package problem0104

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
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				pre: []int{3, 9, 20, 15, 7},
				in:  []int{9, 3, 15, 20, 7},
			},
			a: answer{
				first: 3,
			},
		},
		question{
			p: parameter{
				pre: []int{3, 9, 20, 15, 10, 7},
				in:  []int{9, 3, 10, 15, 20, 7},
			},
			a: answer{
				first: 4,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, maxDepth(ltot(p.pre, p.in)))
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
