package problem0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first []int
	second int
}

type answer struct {
	first int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first: []int{3, 2, 2, 3},
				second: 3,
			},
			a: answer{
				first: 2,
			},
		},
		question{
			p: parameter{
				first: []int{0, 1, 2, 2, 3, 0, 4, 2},
				second: 2,
			},
			a: answer{
				first: 5,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, removeElement(p.first, p.second))
	}
}
