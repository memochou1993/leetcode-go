package problem0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first  string
	second string
}

type answer struct {
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  "hello",
				second: "ll",
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first:  "aaaaa",
				second: "bba",
			},
			a: answer{
				first: -1,
			},
		},
		question{
			p: parameter{
				first:  "bbbbb",
				second: "",
			},
			a: answer{
				first: 0,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, strStr(p.first, p.second))
	}
}
