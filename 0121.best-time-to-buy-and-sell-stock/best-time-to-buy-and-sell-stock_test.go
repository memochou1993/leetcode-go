package problem0121

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
				first: []int{7, 1, 5, 3, 6, 4},
			},
			a: answer{
				first: 5,
			},
		},
		question{
			p: parameter{
				first: []int{7, 6, 4, 3, 1},
			},
			a: answer{
				first: 0,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, maxProfit(p.first))
	}
}
