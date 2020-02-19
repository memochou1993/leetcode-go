package problem0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first  []int
	second []int
}

type answer struct {
	first []int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  []int{1, 2, 4},
				second: []int{1, 3, 4},
			},
			a: answer{
				first: []int{1, 1, 2, 3, 4, 4},
			},
		},
		question{
			p: parameter{
				first:  []int{1, 4},
				second: []int{2, 3},
			},
			a: answer{
				first: []int{1, 2, 3, 4},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, ltos(mergeTwoLists(stol(p.first), stol(p.second))))
	}
}

// convert []int to *ListNode
func stol(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	node := &ListNode{
		Val: nums[0],
	}

	temp := node

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return node
}

// convert *ListNode to []int
func ltos(head *ListNode) []int {
	var node []int

	for head != nil {
		node = append(node, head.Val)
		head = head.Next
	}

	return node
}
