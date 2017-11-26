package hotspot

import (
	"fmt"
	"path"
	"testing"

	"github.com/ready-steady/assert"
)

const (
	fixturePath = "fixtures"
)

func TestNew(t *testing.T) {
	config := prepare("002")
	model := New(config)

	assert.Equal(model.Cores, uint(2), t)
	assert.Equal(model.Nodes, uint(20), t)

	assert.Close(model.C, fixtureC, 2e-15, t)
	assert.Close(model.G, fixtureG, 2e-14, t)
}

func BenchmarkNew(b *testing.B) {
	config := prepare("032")

	for i := 0; i < b.N; i++ {
		New(config)
	}
}

func prepare(floorplan string) *Config {
	return &Config{
		Floorplan:     findFixture(fmt.Sprintf("%s.flp", floorplan)),
		Configuration: findFixture("hotspot.config"),
	}
}

func findFixture(name string) string {
	return path.Join(fixturePath, name)
}
