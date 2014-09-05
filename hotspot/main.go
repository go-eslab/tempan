// Package hotspot provides an interface to HotSpot, a thermal model and
// simulator used in computer-architecture studies.
//
// http://lava.cs.virginia.edu/HotSpot/
package hotspot

// #cgo CFLAGS: -Ibuild
// #cgo LDFLAGS: -Wl,-no_compact_unwind -Lbuild -lhotspot
// #include <stdlib.h>
// #include "hotspot.h"
import "C"

import "unsafe"

type Model struct {
	Cores uint16
	Nodes uint16

	A []float64
	B []float64
	G []float64
}

func Load(floorplan string, config string, params string) *Model {
	cfloorplan := C.CString(floorplan)
	defer C.free(unsafe.Pointer(cfloorplan))

	cconfig := C.CString(config)
	defer C.free(unsafe.Pointer(cconfig))

	cparams := C.CString(params)
	defer C.free(unsafe.Pointer(cparams))

	h := C.newHotSpot(cfloorplan, cconfig, cparams)
	defer C.freeHotSpot(h)

	cc := uint16(h.cores)
	nc := uint32(h.nodes)

	m := &Model{
		Cores: cc,
		Nodes: uint16(nc),

		A: make([]float64, nc),
		B: make([]float64, nc*nc),
		G: make([]float64, nc*nc),
	}

	C.copyA(h, (*C.double)(unsafe.Pointer(&m.A[0])))
	C.copyB(h, (*C.double)(unsafe.Pointer(&m.B[0])))
	C.copyG(h, (*C.double)(unsafe.Pointer(&m.G[0])))

	return m
}
