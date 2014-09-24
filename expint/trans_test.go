package expint

import (
	"testing"

	"github.com/go-math/prob"
	"github.com/go-math/prob/uniform"
	"github.com/go-math/support/assert"
)

func TestComputeTransient(t *testing.T) {
	s, _ := Load(findFixture("002.json"))

	cc := uint32(2)
	sc := uint32(len(fixtureP)) / cc

	Q := make([]float64, cc*sc)
	s.ComputeTransient(fixtureP, Q, nil, sc)

	assert.AlmostEqual(Q, fixtureQ, t)
}

func BenchmarkComputeTransient(b *testing.B) {
	s, _ := Load(findFixture("032.json"))

	cc := uint32(32)
	sc := uint32(1000)
	nc := s.Nodes

	P := prob.Sample(uniform.New(0, 1), cc*sc)
	Q := make([]float64, cc*sc)
	S := make([]float64, nc*sc)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.ComputeTransient(P, Q, S, sc)
	}
}
