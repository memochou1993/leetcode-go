package problem0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first int
}

type answer struct {
	first []int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: 0,
			},
			a: answer{
				first: []int{1},
			},
		},
		question{
			p: parameter{
				first: 2,
			},
			a: answer{
				first: []int{1, 2, 1},
			},
		},
		question{
			p: parameter{
				first: 3,
			},
			a: answer{
				first: []int{1, 3, 3, 1},
			},
		},
		question{
			p: parameter{
				first: 4,
			},
			a: answer{
				first: []int{1, 4, 6, 4, 1},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, getRow(p.first))
	}
}
