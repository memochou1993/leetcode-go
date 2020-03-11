package problem0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p parameter
	a answer
}

type parameter struct {
	first  []int
	second int
	third  []int
	fourth int
}

type answer struct {
	first []int
}

func TestProblem(t *testing.T) {
	questions := []question{
		question{
			p: parameter{
				first:  []int{1, 2, 3, 0, 0, 0},
				second: 3,
				third:  []int{2, 5, 6},
				fourth: 3,
			},
			a: answer{
				first: []int{1, 2, 2, 3, 5, 6},
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		merge(p.first, p.second, p.third, p.fourth)
		assert.Equal(t, a.first, p.first)
	}
}
