package expint

type system struct {
	// D = Cth^(-1/2)
	D []float64

	// A = U * diag(L) * transpose(U)
	U []float64
	L []float64

	B []float64
	C []float64

	E []float64
	F []float64
}
