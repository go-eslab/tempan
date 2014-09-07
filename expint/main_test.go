package expint

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestNew(t *testing.T) {
	s := New(findFixture("problem.json"))

	assert.Equal(s.Cores, uint16(2), t)
	assert.Equal(s.Nodes, uint16(2*4+12), t)
}
