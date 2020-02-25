package problem0038

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
				first: 1,
			},
			a: answer{
				first: "1",
			},
		},
		question{
			p: parameter{
				first: 4,
			},
			a: answer{
				first: "1211",
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, countAndSay(p.first))
	}
}
