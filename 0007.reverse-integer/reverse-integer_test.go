package problem0007

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
				first: 123,
			},
			a: answer{
				first: 321,
			},
		},
		question{
			p: parameter{
				first: -123,
			},
			a: answer{
				first: -321,
			},
		},
		question{
			p: parameter{
				first: 120,
			},
			a: answer{
				first: 21,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, reverse(p.first))
	}
}
