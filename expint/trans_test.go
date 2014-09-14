package expint

import (
	"testing"

	"github.com/go-math/support/assert"
	"github.com/go-math/prob/sample"
)

func TestComputeTransient(t *testing.T) {
	s, _ := Load(findFixture("002.json"))

	cc := uint32(2)
	sc := uint32(len(fixtureP)) / cc

	Q := make([]float64, cc*sc)
	s.ComputeTransient(fixtureP, Q, sc)

	assert.AlmostEqual(Q, fixtureQ, t)
}

func BenchmarkComputeTransient(b *testing.B) {
	s, _ := Load(findFixture("032.json"))

	cc := uint32(32)
	sc := uint32(1000)

	P := sample.Uniform(cc*sc)
	Q := make([]float64, cc*sc)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.ComputeTransient(P, Q, sc)
	}
}
