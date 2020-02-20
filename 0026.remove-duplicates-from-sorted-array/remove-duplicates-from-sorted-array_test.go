package problem0026

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
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: []int{1, 1, 2},
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			a: answer{
				first: 5,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, removeDuplicates(p.first))
	}
}
