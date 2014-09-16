package expint

import (
	"github.com/go-math/linal/matrix"
)

// ComputeTransient performs transient temperature analysis. P is an input
// power profile given as a cc-by-sc matrix where cc is the number of cores,
// and sc is the number of time steps; see TimeStep in Config. Q is the
// corresponding output temperature profile, which is given as a
// cc-by-sc-matrix. S is an optional nc-by-sc matrix, where nc is the number of
// thermal nodes, for the internal usage of the function to prevent repetitive
// memory allocation if the analysis is to be performed several times.
func (s *Self) ComputeTransient(P, Q, S []float64, sc uint32) {
	cc := s.Cores
	nc := s.Nodes

	if S == nil {
		S = make([]float64, nc*sc)
	}

	matrix.Multiply(s.system.F, P, S, nc, cc, sc)

	var i, j, k uint32

	for i, j, k = 1, 0, nc; i < sc; i++ {
		matrix.MultiplyAdd(s.system.E, S[j:k], S[k:k+nc], S[k:k+nc], nc, nc, 1)
		j += nc
		k += nc
	}

	for i = 0; i < cc; i++ {
		for j = 0; j < sc; j++ {
			Q[cc*j+i] = s.system.D[i]*S[nc*j+i] + s.Config.AmbientTemp
		}
	}
}
