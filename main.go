// Package hotspot provides an interface to HotSpot.
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
//               dQ
//     diag(C) * -- + G * (Q - Qamb) = P.
//               dt
//
// The number of thermal nodes is denoted by Nodes. Then, in the above system,
//
//     C is a Nodes-element vector of the thermal capacitance,
//     G is a Nodes-by-Nodes matrix of the thermal conductance,
//     Q is a Nodes-element vector of the heat dissipation,
//     P is a Nodes-element vector of the power consumption, and
//     Qamb is a vector of the ambient temperature.
type Model struct {
	Cores uint
	Nodes uint

	C []float64
	G []float64
}

// New constructs a thermal RC circuit according to the given configuration.
func New(config *Config) *Model {
	floorplan := C.CString(config.Floorplan)
	defer C.free(unsafe.Pointer(floorplan))

	configuration := C.CString(config.Configuration)
	defer C.free(unsafe.Pointer(configuration))

	parameters := C.CString(config.Parameters)
	defer C.free(unsafe.Pointer(parameters))

	hotspot := C.newHotSpot(floorplan, configuration, parameters)
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
