package problem0136

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
				first: []int{2, 2, 1},
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first: []int{4, 1, 2, 1, 2},
			},
			a: answer{
				first: 4,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, singleNumber(p.first))
	}
}
