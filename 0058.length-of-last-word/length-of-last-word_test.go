package problem0058

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
				first: "Hello World",
			},
			a: answer{
				first: 5,
			},
		},
		question{
			p: parameter{
				first: "",
			},
			a: answer{
				first: 0,
			},
		},
		question{
			p: parameter{
				first: " ",
			},
			a: answer{
				first: 0,
			},
		},
		question{
			p: parameter{
				first: "a",
			},
			a: answer{
				first: 1,
			},
		},
		question{
			p: parameter{
				first: " a ",
			},
			a: answer{
				first: 1,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, lengthOfLastWord(p.first))
	}
}
