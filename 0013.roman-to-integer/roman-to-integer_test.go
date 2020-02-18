package problem0013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first string
}

type answer struct {
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: "III",
			},
			a: answer{
				first: 3,
			},
		},
		question{
			p: parameter{
				first: "IV",
			},
			a: answer{
				first: 4,
			},
		},
		question{
			p: parameter{
				first: "IX",
			},
			a: answer{
				first: 9,
			},
		},
		question{
			p: parameter{
				first: "LVIII",
			},
			a: answer{
				first: 58,
			},
		},
		question{
			p: parameter{
				first: "MCMXCIV",
			},
			a: answer{
				first: 1994,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, romanToInt(p.first))
	}
}
