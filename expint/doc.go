// Package expint provides an exponential-integrator-based solver for systems
// of differential-algebraic equations modeling temperature of electronic
// systems.
//
// The initial thermal system is
//
//     C * dQ'/dt + G * (Q' - Qamb) = M * P
//     Q = M**T * Q'
//
// where C and G are the thermal capacitance and conductance matrices,
// respectively; Q' and Q are the temperature vectors of all thermal nodes and
// those that correspond to the processing elements, respectively; Qamb is the
// ambient temperature; P is the power vector of the processing elements; and
// M is a rectangular diagonal matrix whose diagonal elements equal to unity.
//
// The transformed system is
//
//     dS/dt = A * S + B * P
//     Q = B**T * S + Qamb
//
// where
//
//     S = D**(-1) * (Q' - Qamb),
//     A = -D * G * D,
//     B = D * M, and
//     D = C**(-1/2).
//
// The eigendecomposition of A, which is real and symmetric, is
//
//     A = U * diag(L) * U**T.
//
// The solution of the system for a short time interval [0, dt] is based on the
// following recurrence:
//
//     S(t) = E * S(0) + F * P(0)
//
// where
//
//     E = exp(A * dt) = U * diag(exp(li * dt)) * U**T and
//     F = A**(-1) * (exp(A * dt) - I) * B
//       = U * diag((exp(li * dt) - 1) / li) * U**T * B.
package expint
