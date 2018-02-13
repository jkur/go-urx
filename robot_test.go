package urx

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestRobotCreation(t *testing.T) {
	r := NewUR("localhost", "30002")
	pose := r.Getl()
	if pose != mat.NewVecDense(6, []float64{0, 0, 0, 0, 0, 0}) {
		t.Error("Expected 0 found", pose)
	}
}

// this is really bad
func Test_main(t *testing.T) {
	r := NewUR("localhost", "30002")
	r.Run()
}
