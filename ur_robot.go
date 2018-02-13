package urx

import (
	//"bufio"
	"fmt"
	"github.com/lunixbochs/struc"
	"gonum.org/v1/gonum/mat"
	"net"
	"os"
	//	"time"
	"bytes"
	"encoding/binary"
)

type urRobot struct {
	csys mat.VecDense
	addr string
	port string
}

type netPackage struct {
	len  int
	data []byte
}

type UrSecmonHeader struct {
	Length      uint32 `struc:"uint32,big"`
	MessageType uint8  `struc:"uint8"`
}

type UrSecmonPacket struct {
	Header UrSecmonHeader
	Data   []byte
}

type UrSubHeader struct {
	Length      uint32 `struc:"uint32,big"`
	MessageType uint8  `struc:"uint8"`
}

type UrSubPacket struct {
	Header UrSubHeader
	Data   []byte
}

func (s *UrSecmonPacket) String() string {
	return fmt.Sprintf("UrSecmonPacket: Header (Length: %d, Type: %d) - len(Data) := %d", s.Header.Length, s.Header.MessageType, len(s.Data))
}

func (s *UrSubPacket) String() string {
	return fmt.Sprintf("UrSubPacket: Header (Length: %d, Type: %d) - len(Data) := %d", s.Header.Length, s.Header.MessageType, len(s.Data))
}

func (n *netPackage) String() string {
	return fmt.Sprintf("Net: len(%d), %x", n.len, n.data[:n.len])
}

func NewUR() *urRobot {

	robot := new(urRobot)
	return robot
}

func (urrobot *urRobot) getl() *mat.VecDense {
	return mat.NewVecDense(6, []float64{0, 0, 0, 0, 0, 0})
}

func listen(conn *net.TCPConn, c chan *netPackage) {
	//timeoutDuration := 10 * time.Millisecond
	//conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	buf := make([]byte, 2048)
	for {
		reply := &netPackage{}
		size, err := conn.Read(buf)
		if err != nil {
			println("Somthing failed at Read:", err.Error())
		} else {
			//println("Reading data:", size)
			reply.len = size
			reply.data = make([]byte, reply.len)
			copy(reply.data, buf)
			c <- reply
		}
	}
}

func parseUrSecmonPacket(data *netPackage) {
	buf := bytes.NewReader(data.data)
	urh := &UrSecmonHeader{}
	err := struc.Unpack(buf, &urh)
	if err != nil {
		fmt.Println("Something went wrong at parsing secmon packet")
	}
	if urh.MessageType != 16 {
		fmt.Println("Bad packet: Message Type is not 16")
	}
	urp := &UrSecmonPacket{}
	urp.Header = *urh
	urp.Data = make([]byte, urh.Length-5)
	err = binary.Read(buf, binary.BigEndian, &urp.Data)
	if len(urp.Data) != int(urp.Header.Length-5) {
		fmt.Println("Bad packet: Size mismatch")
		return
	}
	fmt.Println(urp)
	parseSubPackages(urp)
}

func parseSubPackages(secmonpkg *UrSecmonPacket) {
	buf := bytes.NewReader(secmonpkg.Data)
	for buf.Len() > 5 {
		urh := &UrSubHeader{}
		err := struc.Unpack(buf, &urh)
		if err != nil {
			fmt.Println("Something went wrong at parsing subpacket")
		}
		switch ptype := urh.MessageType; ptype {
		case ROBOTMODEDATA:
			o := &RobotModeData{}
			if buf.Len() < 47 {
				println("Found Robotmode but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			o.length = urh.Length
			o.ptype = urh.MessageType
			fmt.Println(o)
		case JOINTDATA:
			o := &JointDataContainer{}
			if buf.Len() < 251 {
				println("Found Joint Data but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			o.length = urh.Length
			o.ptype = urh.MessageType
			fmt.Println(o)
		case TOOLDATA:
			o := &ToolData{}
			if buf.Len() < 37 {
				println("Found Tool Data but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			o.length = urh.Length
			o.ptype = urh.MessageType
			fmt.Printf("Tooldata: %+v\n", o)
		case MASTERBOARDDATA:
			// First check length
			if urh.Length == 90 {
				o := &MasterboardDataEuromap{}
				o.length = urh.Length
				o.ptype = urh.MessageType
				if buf.Len() < int(o.length) {
					println("Found Masterboard Data but buffer too small", buf.Len())
					return
				}
				err := struc.Unpack(buf, o)
				if err != nil {
					println("struc error:", err.Error())
				}
				fmt.Printf("Masterboard(Euromap): %+v\n", o)
			} else if urh.Length == 74 {
				o := &MasterboardData{}
				o.length = urh.Length
				o.ptype = urh.MessageType
				if buf.Len() < int(o.length) {
					println("Found Masterboard Data but buffer too small", buf.Len())
					return
				}
				err := struc.Unpack(buf, o)
				if err != nil {
					println("struc error:", err.Error())
				}
				fmt.Printf("Masterboard: %+v\n", o)
			}
		case CARTESIANINFO:
			o := &CartesianInfo{}
			o.length = urh.Length
			o.ptype = urh.MessageType
			if buf.Len() < int(o.length) {
				println("Found CartesianInfob but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			fmt.Printf("CartesianInfo: %+v\n", o)
		case CONFIGURATIONDATA:
			o := &ConfigurationData{}
			o.length = urh.Length
			o.ptype = urh.MessageType
			if buf.Len() < int(o.length) {
				println("Found ConfigurationData but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			fmt.Printf("ConfigurationData: %+v\n", o)
		case ADDITIONALINFO:
			o := &AdditionalInfo{}
			o.length = urh.Length
			o.ptype = urh.MessageType
			if buf.Len() < int(o.length) {
				println("Found AdditionalInfo but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			fmt.Printf("AdditionalInfo: %+v\n", o)
		case FORCEMODEDATA:
			o := &ForceModeData{}
			o.length = urh.Length
			o.ptype = urh.MessageType
			if buf.Len() < int(o.length) {
				println("Found ForceModeData but buffer too small", buf.Len())
				return
			}
			err := struc.Unpack(buf, o)
			if err != nil {
				println("struc error:", err.Error())
			}
			fmt.Printf("ForceModeData: %+v\n", o)

		default:
			fmt.Printf("Unrecognized Packet with type (%d) and Length: %d\n", urh.MessageType, urh.Length)
			if int(urh.Length) < buf.Len() {
				discard_buf := make([]byte, urh.Length-5)
				buf.Read(discard_buf)
			}
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
		c := make(chan *netPackage)
		go listen(conn, c)
		for data := range c {
			//println(data.len, data.data)
			parseUrSecmonPacket(data)
			//fmt.Println(data)
		}
		conn.Close()
	}
}
