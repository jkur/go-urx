package urx

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestRobotCreation(t *testing.T) {
  r := NewUR()
	pose := r.getl()
	if pose != mat.NewVecDense(6, []float64{0, 0, 0, 0, 0, 0}) {
		t.Error("Expected 0 found", pose)
	}
}
