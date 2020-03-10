package problem0083

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
				first: []int{1, 1, 2, 3},
			},
			a: answer{
				first: []int{1, 2, 3},
			},
		},
		question{
			p: parameter{
				first: []int{1, 1, 2, 3, 3},
			},
			a: answer{
				first: []int{1, 2, 3},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, ltos(deleteDuplicates(stol(p.first))))
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
	node := []int{}

	for head != nil {
		node = append(node, head.Val)
		head = head.Next
	}

	return node
}
