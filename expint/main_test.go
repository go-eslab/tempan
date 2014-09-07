package expint

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestNew(t *testing.T) {
	s := New(findFixture("problem.json"))

	assert.Equal(s.Cores, uint16(2), t)
	assert.Equal(s.Nodes, uint16(4*2+12), t)

	assert.Equal(s.system.A, fixtureA, t)
	assert.Equal(s.system.B, fixtureB, t)
	assert.Equal(s.system.C, fixtureC, t)

	assert.Equal(s.system.U, fixtureU, t)
	assert.Equal(s.system.L, fixtureL, t)

	assert.Equal(s.system.E, fixtureE, t)
	assert.Equal(s.system.F, fixtureF, t)
}
