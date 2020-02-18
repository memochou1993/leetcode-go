package problem0020

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
	first bool
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: "()",
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				first: "()[]{}",
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				first: "(]",
			},
			a: answer{
				first: false,
			},
		},
		question{
			p: parameter{
				first: "([)]",
			},
			a: answer{
				first: false,
			},
		},
		question{
			p: parameter{
				first: "{[]}",
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				first: "[",
			},
			a: answer{
				first: false,
			},
		},
		question{
			p: parameter{
				first: "]",
			},
			a: answer{
				first: false,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, isValid(p.first))
	}
}
