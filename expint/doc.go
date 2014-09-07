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
// temperature of those nodes that are active (cores, processing elements).
//
// The transformed system is
//
//     dX/dt = A * X + B * P
//     Q = C * X + Qamb
//
// where
//
//     X = Cth^(1/2) * (Qth - Qamb),
//     A = Cth^(-1/2) * (-Gth) * Cth^(-1/2),
//     B = Cth^(-1/2) * M, and
//     C = B^T.
//
// The eigendecomposition of A, which is real and symmetric, is
//
//     A = U * L * U^(-1) = U * L * U^T = U * L * V.
//
// The solution of the system for a short time interval [0, t] is based on the
// following recurrence:
//
//     X(t) = E * X(0) + F * P(0).
//
// The first coefficient of the recurrence:
//
//     E = exp(A * dt) = U * diag(exp(li * dt)) * V
//
// The second coefficient of the recurrence:
//
//     F = A^(-1) * (exp(A * dt) - I) * B
//       = U * diag((exp(li * dt) - 1) / li) * V * B
//
package expint
