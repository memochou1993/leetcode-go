package problem0070

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
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: 2,
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first: 3,
			},
			a: answer{
				first: 3,
			},
		},
		question{
			p: parameter{
				first: 7,
			},
			a: answer{
				first: 21,
			},
		},
		question{
			p: parameter{
				first: 8,
			},
			a: answer{
				first: 34,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, climbStairs(p.first))
	}
}
