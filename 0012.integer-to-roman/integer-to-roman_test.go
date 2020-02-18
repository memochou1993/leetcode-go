package problem0012

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
	first string
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: 7,
			},
			a: answer{
				first: "VII",
			},
		},
		question{
			p: parameter{
				first: 83,
			},
			a: answer{
				first: "LXXXIII",
			},
		},
		question{
			p: parameter{
				first: 699,
			},
			a: answer{
				first: "DCXCIX",
			},
		},
		question{
			p: parameter{
				first: 1964,
			},
			a: answer{
				first: "MCMLXIV",
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, intToRoman(p.first))
	}
}
