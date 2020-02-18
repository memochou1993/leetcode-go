package problem0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first []string
}

type answer struct {
	first string
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: []string{"flower", "flow", "flight"},
			},
			a: answer{
				first: "fl",
			},
		},
		question{
			p: parameter{
				first: []string{"dog", "racecar", "car"},
			},
			a: answer{
				first: "",
			},
		},
		question{
			p: parameter{
				first: []string{"a"},
			},
			a: answer{
				first: "a",
			},
		},
		question{
			p: parameter{
				first: []string{"aa", "a"},
			},
			a: answer{
				first: "a",
			},
		},
		question{
			p: parameter{
				first: []string{},
			},
			a: answer{
				first: "",
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, longestCommonPrefix(p.first))
	}
}
