package problem0155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-1)

	assert.Equal(t, s.GetMin(), -2)
	assert.Equal(t, s.Top(), -1)

	s.Pop()

	assert.Equal(t, s.Top(), 0)
}
