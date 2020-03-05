package problem0069

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
				first: 0,
			},
			a: answer{
				first: 0,
			},
		},
		question{
			p: parameter{
				first: 1,
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first: 2,
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first: 3,
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first: 4,
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first: 8,
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first: 144,
			},
			a: answer{
				first: 12,
			},
		},
		question{
			p: parameter{
				first: 9000,
			},
			a: answer{
				first: 94,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, mySqrt(p.first))
	}
}
