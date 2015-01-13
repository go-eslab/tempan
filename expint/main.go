package expint

import (
	"math"

	"github.com/ready-steady/linal/decomp"
	"github.com/ready-steady/linal/matrix"
	"github.com/ready-steady/tempan/hotspot"
)

// Solver represents the algorithm for temperature analysis configured for a
// particular problem.
type Solver struct {
	Config Config

	Cores uint32
	Nodes uint32

	system system
}

// New returns an instance of the algorithm set up according to the given
// configuration.
func New(c Config) (*Solver, error) {
	s := &Solver{
		Config: c,
	}

	model := hotspot.New(c.Floorplan, c.HotSpot.Config, c.HotSpot.Params)

	cc := model.Cores
	nc := model.Nodes

	var i, j uint32

	// Reusing model.G to store A and model.C to store D.
	A := model.G
	D := model.C
	for i = 0; i < nc; i++ {
		D[i] = math.Sqrt(1 / model.C[i])
	}
	for i = 0; i < nc; i++ {
		for j = 0; j < nc; j++ {
			A[j*nc+i] = -1 * D[i] * D[j] * A[j*nc+i]
		}
	}

	// Reusing A (which is model.G) to store U.
	U := A
	Λ := make([]float64, nc)
	if err := decomp.SymEig(A, U, Λ, nc); err != nil {
		return nil, err
	}

	Δt := c.TimeStep

	coef := make([]float64, nc)
	temp := make([]float64, nc*nc)

	for i = 0; i < nc; i++ {
		coef[i] = math.Exp(Δt * Λ[i])
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
		coef[i] = (coef[i] - 1) / Λ[i]
	}
	for i = 0; i < nc; i++ {
		for j = 0; j < cc; j++ {
			temp[j*nc+i] = coef[i] * U[i*nc+j] * D[j]
		}
	}

	F := make([]float64, nc*cc)
	matrix.Multiply(U, temp, F, nc, nc, cc)

	s.Cores = model.Cores
	s.Nodes = model.Nodes

	s.system.D = D

	s.system.Λ = Λ
	s.system.U = U

	s.system.E = E
	s.system.F = F

	return s, nil
}

// Load returns an instance of the algorithm set up according to the given
// configuration file.
func Load(path string) (*Solver, error) {
	config, err := loadConfig(path)
	if err != nil {
		return nil, err
	}

	return New(config)
}
