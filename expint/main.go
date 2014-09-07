package expint

import (
	"github.com/go-eslab/tempan/hotspot"
)

// Solver represents the algorithm configured for a particular platform.
type Solver struct {
	Config Config

	Cores  uint16
	Nodes  uint16

	system system
}

// New returns an instance of the solver configured according to the given
// arguments.
func New(configPath string) *Solver {
	solver := new(Solver)

	c := &solver.Config
	c.load(configPath)

	h := hotspot.New(c.Floorplan, c.HotSpot.Config, c.HotSpot.Params)

	cc := h.Cores
	nc := h.Nodes

	solver.Cores = cc
	solver.Nodes = nc

	return solver
}
