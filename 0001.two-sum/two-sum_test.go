package problem0001

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
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
	first []int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  []int{2, 7, 11, 15},
				second: 9,
			},
			a: answer{
				first: []int{0, 1},
			},
		},
		question{
			p: parameter{
				first:  []int{2, 7, 11, 15},
				second: 8,
			},
			a: answer{
				first: []int{},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, secondSum(p.first, p.second), a.first)
	}
}
