package hotspot

import (
	"fmt"
	"path"
	"testing"

	"github.com/ready-steady/support/assert"
)

const (
	fixturePath = "fixtures"
)

func TestNew(t *testing.T) {
	config := prepare("002", "")
	model := New(config)

	assert.Equal(model.Cores, uint(2), t)
	assert.Equal(model.Nodes, uint(20), t)

	assert.EqualWithin(model.C, fixtureC, 2e-15, t)
	assert.EqualWithin(model.G, fixtureG, 2e-14, t)
}

func TestNewWithParams(t *testing.T) {
	config := prepare("002", "t_chip 0.00042 k_chip 42")
	model := New(config)

	assert.EqualWithin(model.C, fixtureC42, 1e-15, t)
	assert.EqualWithin(model.G, fixtureG42, 1e-15, t)
}

func BenchmarkNew(b *testing.B) {
	config := prepare("032", "")

	for i := 0; i < b.N; i++ {
		New(config)
	}
}

func prepare(floorplan string, parameters string) *Config {
	return &Config{
		Floorplan:     findFixture(fmt.Sprintf("%s.flp", floorplan)),
		Configuration: findFixture("hotspot.config"),
		Parameters:    parameters,
	}
}

func findFixture(name string) string {
	return path.Join(fixturePath, name)
}
