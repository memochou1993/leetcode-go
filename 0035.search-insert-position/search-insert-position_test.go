package problem0035

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
	second int
}

type answer struct {
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  []int{1, 3, 5, 6},
				second: 5,
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first:  []int{1, 3, 5, 6},
				second: 2,
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first:  []int{1, 3, 5, 6},
				second: 7,
			},
			a: answer{
				first: 4,
			},
		},
		question{
			p: parameter{
				first:  []int{1, 3, 5, 6},
				second: 0,
			},
			a: answer{
				first: 0,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, searchInsert(p.first, p.second))
	}
}
