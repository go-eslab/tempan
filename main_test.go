package hotspot

import (
	"path"
	"testing"

	"github.com/ready-steady/support/assert"
)

const (
	fixturePath = "fixtures"
)

func TestNew(t *testing.T) {
	model := New(findFixture("002.flp"), findFixture("hotspot.config"), "")

	assert.Equal(model.Cores, uint(2), t)
	assert.Equal(model.Nodes, uint(20), t)

	assert.EqualWithin(model.C, fixtureC, 2e-15, t)
	assert.EqualWithin(model.G, fixtureG, 2e-14, t)
}

func TestNewWithParams(t *testing.T) {
	model := New(findFixture("002.flp"), findFixture("hotspot.config"),
		"t_chip 0.00042 k_chip 42")

	assert.EqualWithin(model.C, fixtureC42, 1e-15, t)
	assert.EqualWithin(model.G, fixtureG42, 1e-15, t)
}

func BenchmarkLoad(b *testing.B) {
	floorplan := findFixture("032.flp")
	config := findFixture("hotspot.config")
	params := ""

	for i := 0; i < b.N; i++ {
		New(floorplan, config, params)
	}
}

func findFixture(name string) string {
	return path.Join(fixturePath, name)
}
