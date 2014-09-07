package expint

import (
	"github.com/go-math/linal/matrix"
)

type system struct {
	A *matrix.Matrix
	B *matrix.Matrix
	C *matrix.Matrix

	U *matrix.Matrix
	L *matrix.Matrix
	V *matrix.Matrix

	E *matrix.Matrix
	F *matrix.Matrix
}
