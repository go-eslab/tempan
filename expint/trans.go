package expint

import (
	"github.com/go-math/linal/matrix"
)

// ComputeTransient performs transient temperature analysis. P is an input
// power profile given as a cc-by-sc matrix where cc is the number of cores
// (processing elements), and sc is the number of time steps (power samples);
// see TimeStep in Config. Q is the corresponding output temperature profile,
// which is given as a cc-by-sc-matrix.
func (s *Solver) ComputeTransient(P, Q []float64, cc uint32, sc uint32) {
	nc := s.Nodes

	X := make([]float64, nc*sc)
	matrix.Multiply(s.system.F, P, X, nc, cc, sc)

	var i, j, k uint32

	for i, j, k = 1, 0, nc; i < sc; i++ {
		matrix.MultiplyAdd(s.system.E, X[j:k], X[k:k+nc], X[k:k+nc], nc, nc, 1)
		j += nc
		k += nc
	}

	for i = 0; i < cc; i++ {
		for j = 0; j < sc; j++ {
			Q[cc*j+i] = s.system.D[i] * X[nc*j+i] + s.Config.AmbientTemp
		}
	}
}
