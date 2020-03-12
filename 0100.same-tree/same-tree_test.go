package problem0100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	firstPre  []int
	firstIn   []int
	secondPre []int
	secondIn  []int
}

type answer struct {
	first bool
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				firstPre:  []int{1, 2, 3},
				firstIn:   []int{2, 1, 3},
				secondPre: []int{1, 2, 3},
				secondIn:  []int{2, 1, 3},
			},
			a: answer{
				first: true,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, isSameTree(ltot(p.firstPre, p.firstIn), ltot(p.secondPre, p.secondIn)))
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
