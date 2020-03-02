package problem0066

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
				first: []int{1, 9, 9},
			},
			a: answer{
				first: []int{2, 0, 0},
			},
		},
		question{
			p: parameter{
				first: []int{4, 3, 2, 1},
			},
			a: answer{
				first: []int{4, 3, 2, 2},
			},
		},
		question{
			p: parameter{
				first: []int{9},
			},
			a: answer{
				first: []int{1, 0},
			},
		},
		question{
			p: parameter{
				first: []int{5, 8, 7, 8, 8},
			},
			a: answer{
				first: []int{5, 8, 7, 8, 9},
			},
		},
		question{
			p: parameter{
				first: []int{2, 4, 9, 3, 9},
			},
			a: answer{
				first: []int{2, 4, 9, 4, 0},
			},
		},
		question{
			p: parameter{
				first: []int{8, 9, 9, 9},
			},
			a: answer{
				first: []int{9, 0, 0, 0},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, plusOne(p.first))
	}
}
