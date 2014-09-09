package expint

import (
	"math"

	"github.com/go-eslab/tempan/hotspot"
	"github.com/go-math/linal/decomp"
	"github.com/go-math/linal/matrix"
)

// Solver represents the algorithm for temperature analysis configured for a
// particular problem.
type Solver struct {
	Config *Config

	Cores uint32
	Nodes uint32

	system system
}

// New returns an instance of the solver set up according to the given
// configuration.
func New(c *Config) (*Solver, error) {
	s := &Solver{
		Config: c,
	}

	h := hotspot.New(c.Floorplan, c.HotSpot.Config, c.HotSpot.Params)

	cc := h.Cores
	nc := h.Nodes

	var i, j uint32

	// Reusing h.G to store A and h.C to store D.
	A := h.G
	D := h.C
	for i = 0; i < nc; i++ {
		D[i] = math.Sqrt(1 / h.C[i])
	}
	for i = 0; i < nc; i++ {
		for j = 0; j < nc; j++ {
			A[j*nc+i] = -1 * D[i] * D[j] * A[j*nc+i]
		}
	}

	// Reusing A (which is h.G) to store U.
	U := A
	L := make([]float64, nc)
	if err := decomp.SymEig(A, U, L, nc); err != nil {
		return nil, err
	}

	dt := c.TimeStep

	coef := make([]float64, nc)
	temp := make([]float64, nc*nc)

	for i = 0; i < nc; i++ {
		coef[i] = math.Exp(dt * L[i])
	}
	for i = 0; i < nc; i++ {
		for j = 0; j < nc; j++ {
			temp[j*nc+i] = coef[i] * U[i*nc+j]
		}
	}

	E := make([]float64, nc*nc)
	matrix.Multiply(U, temp, E, nc, nc, nc)

	// Technically, temp = temp[0 : nc*cc].
	for i = 0; i < nc; i++ {
		coef[i] = (coef[i] - 1) / L[i]
	}
	for i = 0; i < nc; i++ {
		for j = 0; j < cc; j++ {
			temp[j*nc+i] = coef[i] * U[i*nc+j] * D[j]
		}
	}

	F := make([]float64, nc*cc)
	matrix.Multiply(U, temp, F, nc, nc, cc)

	s.Cores = h.Cores
	s.Nodes = h.Nodes

	s.system.D = D

	s.system.L = L
	s.system.U = U

	s.system.E = E
	s.system.F = F

	return s, nil
}

// Load returns an instance of the solver set up according to the given
// configuration file.
func Load(path string) (*Solver, error) {
	config := new(Config)

	if err := config.load(path); err != nil {
		return nil, err
	}

	return New(config)
}
