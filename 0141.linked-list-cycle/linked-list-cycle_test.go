package problem0141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	nums []int
	pos  int
}

type answer struct {
	first bool
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				nums: []int{3, 2, 0, -4},
				pos:  1,
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				nums: []int{1, 2},
				pos:  0,
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				nums: []int{1},
				pos:  -1,
			},
			a: answer{
				first: false,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, hasCycle(stolWithCycle(p.nums, p.pos)))
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

// convert []int to *ListNode
func stolWithCycle(nums []int, pos int) *ListNode {
	head := stol(nums)

	if pos == -1 {
		return head
	}

	c := head

	for pos > 0 {
		c = c.Next
		pos--
	}

	tail := c

	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = c

	return head
}
