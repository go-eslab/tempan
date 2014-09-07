package expint

import (
	"math"
	"testing"

	"github.com/go-math/support/assert"
)

func TestNew(t *testing.T) {
	s, _ := New(findFixture("problem.json"))

	assert.Equal(s.Cores, uint16(2), t)
	assert.Equal(s.Nodes, uint16(4*2+12), t)

	assert.AlmostEqual(s.system.D, fixtureD, t)

	assert.AlmostEqual(abs(s.system.U), abs(fixtureU), t)
	assert.AlmostEqual(s.system.L, fixtureL, t)

	assert.AlmostEqual(s.system.E, fixtureE, t)
	assert.AlmostEqual(s.system.F, fixtureF, t)
}

func abs(A []float64) []float64 {
	B := make([]float64, len(A))

	for i := range B {
		B[i] = math.Abs(A[i])
	}

	return B
}
