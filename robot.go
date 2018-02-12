package urx

import (
	"gonum.org/v1/gonum/mat"
)

type Robot interface {
	// Reference Coordinate
	getl() mat.VecDense
	movej(pose mat.VecDense, acc, vel float64)
	movel(pose mat.VecDense, acc, vel float64)
}
