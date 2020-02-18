package problem0009

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
	first bool
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: 121,
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				first: -121,
			},
			a: answer{
				first: false,
			},
		},
		question{
			p: parameter{
				first: 10,
			},
			a: answer{
				first: false,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, isPalindrome(p.first))
	}
}
