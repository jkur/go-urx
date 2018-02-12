package urx

import (
	//"bufio"
	"encoding/binary"
	"fmt"
	"github.com/lunixbochs/struc"
	"gonum.org/v1/gonum/mat"
	//	"io"
	"net"
	"os"
	//	"time"
	"bytes"
)

type urRobot struct {
	csys mat.VecDense
	addr string
	port string
}

type netPackage struct {
	len  uint32
	data []byte
}

type RobotModeData struct {
	Length                   uint32  `struc:"uint32,big"`
	Ptype                    int     `struc:"uint8"`
	Timestamp                uint64  `struc:"uint64,big"`
	PhysicalRobotConnected   bool    `struc:"bool"`
	RealRobotEnabled         bool    `struc:"bool"`
	RobotPowerOn             bool    `struc:"bool"`
	EmergencyStopped         bool    `struc:"bool"`
	ProtectiveStopped        bool    `struc:"bool"`
	ProgramRunning           bool    `struc:"bool"`
	ProgramPaused            bool    `struc:"bool"`
	RobotMode                uint8   `struc:"uint8"`
	ControlMode              uint8   `struc:"uint8"`
	TargetSpeedFraction      float64 `struc:"float64"`
	SpeedScaling             float64 `struc:"float64"`
	TargetSpeedFractionLimit float64 `struc:"float64"`
	Reserved                 uint8   `struc:"uint8"`
}

func (r *RobotModeData) String() string {
	return fmt.Sprintf("RobotModeData (%d): PhysicalRobotConnected (%t), RealRobotEnabled (%t), RobotPowerOn (%t), EmergencyStopped (%t), ProtectiveStopped (%t), ProgrammRunning (%t), ProgramPaused (%t), RobotMode (%d), ControlMode (%d), SpeedScaling (%f)",
		r.Timestamp, r.PhysicalRobotConnected, r.RealRobotEnabled, r.RobotPowerOn, r.EmergencyStopped, r.EmergencyStopped, r.ProgramRunning, r.ProgramPaused, r.RobotMode, r.ControlMode, r.TargetSpeedFraction)
}

type urMessage struct {
	len   uint32
	mtype byte
	data  []byte
}

func (m urMessage) String() string {
	return fmt.Sprintf("RobotMessage %d with length: %d", m.mtype, m.len)
}

func NewUR() *urRobot {

	robot := new(urRobot)
	return robot
}

func (urrobot *urRobot) getl() *mat.VecDense {
	return mat.NewVecDense(6, []float64{0, 0, 0, 0, 0, 0})
}

func listen(conn *net.TCPConn, c chan netPackage) {
	//timeoutDuration := 10 * time.Millisecond
	//conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	for {
		reply := new(netPackage)
		reply.data = make([]byte, 2048)
		size, err := conn.Read(reply.data)
		if err != nil {
			println("Somthing failed at Read:", err.Error())
		} else {
			println("Readdind data:", size)
			reply.len = uint32(size)
			c <- *reply
		}
	}
}

func parseMessage(data *netPackage) {
	// validity check
	packet_length := binary.BigEndian.Uint32(data.data[0:4])
	mtype := data.data[4]
	if packet_length != data.len {
		println("broken packet. aborting")
		println("length: ", packet_length, data.len)
	} else {
		println("Message type: ", mtype)
	}
	parseSubPackage(data.data[5:], uint32(len(data.data)-5))
}

func parseSubPackage(data []byte, size uint32) {
	if size > 5 {
		packet_length := binary.BigEndian.Uint32(data[0:4])
		mtype := data[4]
		if packet_length > size {
			println(" packet too long")
		} else {
			println("packet length", packet_length, len(data))
			println("SubType:", mtype)
		}
		if packet_length < size {
			//parseSubPackage(data[packet_length:], size-packet_length)
		}
		if mtype == 0 {
			o := &RobotModeData{}
			err := struc.Unpack(bytes.NewReader(data), o)
			if err != nil {
				println("struc error")
			}
			fmt.Println(o)
		}
	}
}

func (urrobot *urRobot) Start(addr, port string) {
	urrobot.addr = addr
	urrobot.port = port
	tcpAddr, err := net.ResolveTCPAddr("tcp", urrobot.addr+":"+urrobot.port)
	for {
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}
		c := make(chan netPackage)
		go listen(conn, c)
		for data := range c {
			println(data.len, data.data)
			parseMessage(&data)
		}
		conn.Close()
	}
}
