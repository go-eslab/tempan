package hotspot

import (
	"path"
	"testing"

	"github.com/goesd/support/assert"
)

const (
	fixturePath = "fixtures"
)

func TestLoad(t *testing.T) {
	model := Load(findFixture("002.flp"), findFixture("hotspot.config"), "")

	assert.Equal(model.Cores, uint16(2), t)
	assert.Equal(model.Nodes, uint16(20), t)

	assert.AlmostEqual(model.A, fixtureA, t)
	assert.AlmostEqual(model.B, fixtureB, t)
	assert.AlmostEqual(model.G, fixtureG, t)
}

func findFixture(name string) string {
	return path.Join(fixturePath, name)
}
