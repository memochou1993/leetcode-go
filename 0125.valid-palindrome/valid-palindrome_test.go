package problem0125

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
				first: "A man, a plan, a canal: Panama",
			},
			a: answer{
				first: true,
			},
		},
		question{
			p: parameter{
				first: "0P",
			},
			a: answer{
				first: false,
			},
		},
		question{
			p: parameter{
				first: "race a car",
			},
			a: answer{
				first: false,
			},
		},
	}

	for _, q := range questions {
		a, p := q.a, q.p
		assert.Equal(t, a.first, isPalindrome(p.first))
	}
}
