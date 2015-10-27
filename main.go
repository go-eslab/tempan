// Package hotspot provides an interface to HotSpot.
//
// http://lava.cs.virginia.edu/HotSpot
package hotspot

// #cgo CFLAGS: -Isource -Isource/hotspot
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include <string.h>
// #include <circuit.h>
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
	const (
		sizeOfFloat64 = 8
	)

	floorplan := C.CString(config.Floorplan)
	defer C.free(unsafe.Pointer(floorplan))

	configuration := C.CString(config.Configuration)
	defer C.free(unsafe.Pointer(configuration))

	circuit := C.newCircuit(floorplan, configuration)
	defer C.dropCircuit(circuit)

	cc := uint(circuit.units)
	nc := uint(circuit.nodes)

	m := &Model{
		Cores: cc,
		Nodes: nc,

		C: make([]float64, nc),
		G: make([]float64, nc*nc),
	}

	C.memcpy(unsafe.Pointer(&m.C[0]), unsafe.Pointer(circuit.capacitance),
		C.size_t(sizeOfFloat64*nc))
	C.memcpy(unsafe.Pointer(&m.G[0]), unsafe.Pointer(circuit.conductance),
		C.size_t(sizeOfFloat64*nc*nc))

	return m
}
