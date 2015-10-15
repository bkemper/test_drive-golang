// SEE: https://golang.org/doc/code.html#Testing
package basics_test

import (
        "testing"
        "./../basics"
)

func TestDivide(t *testing.T) {
  quotient, remainder := basics.Divide(101,10)

  if quotient != 10 {
    t.Errorf("Expected quotient to equal 10 except received %v", quotient)
  }

  if remainder != 1 {
    t.Errorf("Expected remainder to equal 1 except received %v", remainder)
  }
}

func TestDivideWithZero(t *testing.T) {
  quotient, remainder := basics.Divide(100,0)

  if quotient != 0 && remainder != 0 {
    t.Error("Expected both quotient and remainder to be zero when dividing by 0")
  }
}
