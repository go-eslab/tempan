// Package hotspot constructs thermal RC circuits for multiprocessor systems
// based on the block model of HotSpot.
//
// http://lava.cs.virginia.edu/HotSpot/
package hotspot

// #cgo CFLAGS: -Ihotspot
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "hotspot.h"
import "C"

import "unsafe"

// Model represents the block variant of the HotSpot model. The thermal system
// under consideration is as follows:
//
//     diag(C) * dT/dt + G * (T - Tamb) = P
//
// where
//
//     Cores is the number of cores (active thermal nodes),
//     Nodes is the number of thermal nodes (4 * Cores + 12),
//
//     T is a Nodes-element vector of temperature,
//     P is a Cores-element vector of power,
//     C is a Nodes-element vector of capacitance, and
//     G is a (Nodes x Nodes) matrix of conductance.
type Model struct {
	Cores uint
	Nodes uint

	C []float64
	G []float64
}

// New returns the block HotSpot model corresponding to the given floorplan
// file, configuration file, and parameter line. The parameter line bears the
// same meaning as the command-line arguments of the HotSpot tool. The names of
// parameters should not include dashes in front of them; for instance, params
// can be "t_chip 0.00015 k_chip 100.0".
func New(floorplan string, config string, params string) *Model {
	cfloorplan := C.CString(floorplan)
	defer C.free(unsafe.Pointer(cfloorplan))

	cconfig := C.CString(config)
	defer C.free(unsafe.Pointer(cconfig))

	cparams := C.CString(params)
	defer C.free(unsafe.Pointer(cparams))

	hotspot := C.newHotSpot(cfloorplan, cconfig, cparams)
	defer C.freeHotSpot(hotspot)

	cc := uint(hotspot.cores)
	nc := uint(hotspot.nodes)

	m := &Model{
		Cores: cc,
		Nodes: nc,

		C: make([]float64, nc),
		G: make([]float64, nc*nc),
	}

	C.copyC(hotspot, (*C.double)(unsafe.Pointer(&m.C[0])))
	C.copyG(hotspot, (*C.double)(unsafe.Pointer(&m.G[0])))

	return m
}
