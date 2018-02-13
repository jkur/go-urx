package urx

import (
	"gonum.org/v1/gonum/mat"
)

type urRobot struct {
	Csys   mat.VecDense
	Addr   string
	Port   string
	secmon *Ur_Secondary_Monitor
}

func NewUR(addr, port string) *urRobot {

	robot := new(urRobot)
	robot.Addr = addr
	robot.Port = port
	robot.secmon = NewUr_Secondary_Monitor(robot)
	return robot
}
func (urrobot *urRobot) Run() {
	urrobot.secmon.Start()
}

func (urrobot *urRobot) Getl() *mat.VecDense {
	return mat.NewVecDense(6, []float64{0, 0, 0, 0, 0, 0})
}
