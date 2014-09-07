package expint

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestConfigLoad(t *testing.T) {
	c := new(Config)
	err := c.load(findFixture("problem.json"))

	assert.Success(err, t)

	assert.Equal(c.Floorplan, findFixture("002.flp"), t)
	assert.Equal(c.HotSpot.Config, findFixture("hotspot.config"), t)
	assert.Equal(c.HotSpot.Params, "", t)
	assert.Equal(c.TimeStep, 1e-3, t)
}
