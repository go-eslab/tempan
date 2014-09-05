package hotspot

import (
	"path"
	"testing"

	"github.com/go-eslab/support/assert"
)

const (
	fixturePath = "fixtures"
)

func TestLoad(t *testing.T) {
	model := Load(findFixture("002.flp"), findFixture("hotspot.config"), "")

	assert.Equal(model.Cores, uint16(2), t)
	assert.Equal(model.Nodes, uint16(20), t)

	assert.AlmostEqual(model.C, fixtureC, t)
	assert.AlmostEqual(model.G, fixtureG, t)
}

func BenchmarkLoad(b *testing.B) {
	floorplan := findFixture("002.flp")
	config := findFixture("hotspot.config")
	params := ""

	for i := 0; i < b.N; i++ {
		Load(floorplan, config, params)
	}
}

func findFixture(name string) string {
	return path.Join(fixturePath, name)
}
