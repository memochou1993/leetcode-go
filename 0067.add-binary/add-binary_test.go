package problem0067

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
	first string
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  "11",
				second: "1",
			},
			a: answer{
				first: "100",
			},
		},
		question{
			p: parameter{
				first:  "1010",
				second: "1011",
			},
			a: answer{
				first: "10101",
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, addBinary(p.first, p.second))
	}
}
