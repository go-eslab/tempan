package expint

type system struct {
	// D = C**(1/2)
	D []float64

	// A = U * diag(L) * U**T
	U []float64
	L []float64

	// E = exp(A * dt) = U * diag(exp(li * dt)) * U**T
	E []float64

	// F = A**(-1) * (exp(A * dt) - I) * B
	//   = U * diag((exp(li * dt) - 1) / li) * U**T * B
	F []float64
}
