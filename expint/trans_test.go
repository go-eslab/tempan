package expint

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestSolverComputeTransient(t *testing.T) {
	s, _ := New(findFixture("002.json"))

	cc := uint32(2)
	sc := uint32(len(fixtureP)) / cc

	Q := make([]float64, cc*sc)
	s.ComputeTransient(fixtureP, Q, cc, sc)

	assert.AlmostEqual(Q, fixtureQ, t)
}
