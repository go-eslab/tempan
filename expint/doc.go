// Package expint provides an exponential-integrator-based solver of systems
// of differential-algebraic equations modeling temperature of electronic
// systems.
//
// The initial thermal system is
//
//     Cth * dQth/dt + Gth * (Qth - Qamb) = M * P
//     Q = M^T * Qth
//
// where Qth is the temperature of all thermal nodes while Q is the
// temperature of those nodes that are active (processing elements).
//
// The transformed system is
//
//     dX/dt = A * X + B * P
//     Q = C * X + Qamb
//
// where
//
//     X = D^(-1) * (Qth - Qamb),
//     A = D * (-Gth) * D,
//     B = D * M,
//     C = B^T, and
//     D = Cth^(-1/2).
//
// The eigendecomposition of A, which is real and symmetric, is
//
//     A = U * diag(L) * U^(-1) = U * diag(L) * U^T.
//
// The solution of the system for a short time interval [0, dt] is based on the
// following recurrence:
//
//     X(t) = E * X(0) + F * P(0).
//
// The first coefficient of the recurrence:
//
//     E = exp(A * dt) = U * diag(exp(li * dt)) * U^T.
//
// The second coefficient of the recurrence:
//
//     F = A^(-1) * (exp(A * dt) - I) * B
//       = U * diag((exp(li * dt) - 1) / li) * U^T * B.
package expint
