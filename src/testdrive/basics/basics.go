// Package basics is a compilation of everything you learn from
// http://tour.golang.org/basics smashed into a single package
package basics

// Factored import package declaration
// SEE: http://tour.golang.org/basics/2
import (
        "math"
)

// Constants declaration; cannot be declared with :=
// SEE: http://tour.golang.org/basics/11 (variable types)
// SEE: http://tour.golang.org/basics/14 (type inference)
// SEE: http://tour.golang.org/basics/15 (constants)
// SEE: http://tour.golang.org/basics/16 (numeric constants)
const (
  Min = 0
  Max = 1 << 100
)

// Package variables
// SEE: http://tour.golang.org/basics/8 (declaring variables)
// SEE: http://tour.golang.org/basics/9 (initializing variables)
// SEE: http://tour.golang.org/basics/10 (short declaration)
var basic = true

// isZero is a package function to determine if a float value is zero
// SEE: http://tour.golang.org/basics/4 (functions)
func isZero(value float64) bool {
  // Zero value assignment; also works var zero = 0, zero := 0, zero := 0.0
  // SEE: http://tour.golang.org/basics/12
  var zero float64

  return value == zero
}

// Divide computes the number of times the denominator fits in the numerator
// SEE: http://tour.golang.org/basics/3 (exportable function)
// SEE: http://tour.golang.org/basics/5 (shorthand declaration)
// SEE: http://tour.golang.org/basics/6 (multiple return values)
func Divide(numerator, denominator float64) (quotient, remainder int) {

  // Guard against dividing by zero
  if isZero(denominator) { return }

  // SEE: http://tour.golang.org/basics/13 (type conversion)
  quotient = int(numerator / denominator)
  remainder = int(math.Remainder(numerator, denominator))

  // SEE: http://tour.golang.org/basics/7 (named or "naked" return)
  return
}
