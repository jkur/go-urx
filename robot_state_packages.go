package urx

import (
	"fmt"
	//"github.com/lunixbochs/struc"
)

const (
	ROBOTMODEDATA = iota
	JOINTDATA
	TOOLDATA
	MASTERBOARDDATA
	CARTESIANINFO
	KINEMATICSINFO
	CONFIGURATIONDATA
	FORCEMODEDATA
	ADDITIONALINFO
	CALIBRATIONDATA
	SAFETYDATA
)

type RobotModeData struct {
	length                   uint32
	ptype                    uint8
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
	TargetSpeedFraction      float64 `struc:"float64,big"`
	SpeedScaling             float64 `struc:"float64,big"`
	TargetSpeedFractionLimit float64 `struc:"float64,big"`
	Reserved                 uint8   `struc:"uint8"`
}

func (r *RobotModeData) String() string {
	return fmt.Sprintf("RobotModeData (%d): PhysicalRobotConnected (%t), RealRobotEnabled (%t), RobotPowerOn (%t), EmergencyStopped (%t), ProtectiveStopped (%t), ProgrammRunning (%t), ProgramPaused (%t), RobotMode (%d), ControlMode (%d), TargetSpeedFraction (%f)",
		r.Timestamp, r.PhysicalRobotConnected, r.RealRobotEnabled, r.RobotPowerOn, r.EmergencyStopped, r.EmergencyStopped, r.ProgramRunning, r.ProgramPaused, r.RobotMode, r.ControlMode, r.TargetSpeedFraction)
}

type JointDataContainer struct { // 251 bytes
	length uint32
	ptype  uint8
	Joint0 JointData `struc:"JointData"`
	Joint1 JointData `struc:"JointData"`
	Joint2 JointData `struc:"JointData"`
	Joint3 JointData `struc:"JointData"`
	Joint4 JointData `struc:"JointData"`
	Joint5 JointData `struc:"JointData"`
}

func (j *JointDataContainer) String() string {
	return fmt.Sprintf("JointDataContainer (%d): (%+v)(%+v)(%+v)(%+v)(%+v)(%+v)", j.length, j.Joint0, j.Joint1, j.Joint2, j.Joint3, j.Joint4, j.Joint5)
}

type JointData struct {
	QActual   float64 `struc:"float64,big"`
	QTarget   float64 `struc:"float64,big"`
	QDActual  float64 `struc:"float64,big"`
	IActual   float32 `struc:"float32"`
	VActual   float32 `struc:"float32"`
	TMotor    float32 `struc:"float32"`
	Unused    float32 `struc:"float32"`
	JointMode uint8   `struc:"uint8"`
}

type ToolData struct { // 37 bytes
	length            uint32
	ptype             uint8
	AnalogInputRange2 int8    `struc:"int8"`
	AnalogInputRange3 int8    `struc:"int8"`
	AnalogInput2      float64 `struc:"float64"`
	AnalogInput3      float64 `struc:"float64"`
	ToolVoltage       float32 `struc:"float32"`
	ToolOutputVoltage uint8   `struc:"uint8"`
	ToolCurrent       float32 `struc:"float32"`
	ToolTemperature   float32 `struc:"float32"`
	ToolMode          uint8   `struc:"uint8"`
}

//func (t *ToolData) String() string {
//	return fmt.Sprintf("ToolData: (%+v)", t)
//}

type MasterboardData struct {
	length                   uint32
	ptype                    uint8
	DigitalInputBits         int32   `struc:"int32"`
	DigitalOutputBits        int32   `struc:"int32"`
	AnalogInputRange0        int8    `struc:"int8"`
	AnalogInputRange1        int8    `struc:"int8"`
	AnalogInput0             float64 `struc:"float64"`
	AnalogInput1             float64 `struc:"float64"`
	AnalogOutputDomain0      int8    `struc:"int8"`
	AnalogOutputDomain1      int8    `struc:"int8"`
	AnalogOutput0            float64 `struc:"float64"`
	AnalogOutput1            float64 `struc:"float64"`
	MasterboardTemperature   float32 `struc:"float32"`
	RobotVoltage             float32 `struc:"float32"`
	RobotCurrent             float32 `struc:"float32"`
	MasterIOCurrent          float32 `struc:"float32"`
	SafetyMode               uint8   `struc:"uint8"`
	ReducedMode              uint8   `struc:"uint8"`
	Euromap67                int8    `struc:"int8"`
	Reserved                 uint32  `struc:"uint32"`
	OperationalMode          uint8   `struc:"uint8"`
	ThreePositionDeviceInput uint8   `struc:"uint8"`
}

type MasterboardDataEuromap struct {
	length                   uint32
	ptype                    uint8
	DigitalInputBits         int32   `struc:"int32"`
	DigitalOutputBits        int32   `struc:"int32"`
	AnalogInputRange0        int8    `struc:"int8"`
	AnalogInputRange1        int8    `struc:"int8"`
	AnalogInput0             float64 `struc:"float64"`
	AnalogInput1             float64 `struc:"float64"`
	AnalogOutputDomain0      int8    `struc:"int8"`
	AnalogOutputDomain1      int8    `struc:"int8"`
	AnalogOutput0            float64 `struc:"float64"`
	AnalogOutput1            float64 `struc:"float64"`
	MasterboardTemperature   float32 `struc:"float32"`
	RobotVoltage             float32 `struc:"float32"`
	RobotCurrent             float32 `struc:"float32"`
	MasterIOCurrent          float32 `struc:"float32"`
	SafetyMode               uint8   `struc:"uint8"`
	ReducedMode              uint8   `struc:"uint8"`
	Euromap67                int8    `struc:"int8"`
	EuromapInputBits         int32   `struc:"int32"`
	EuromapOutputBits        int32   `struc:"int32"`
	EuromapVoltage           float64 `struc:"float64"`
	EuromapCurrent           float64 `struc:"float64"`
	Reserved                 uint32  `struc:"uint32"`
	OperationalMode          uint8   `struc:"uint8"`
	ThreePositionDeviceInput uint8   `struc:"uint8"`
}

type CartesianInfo struct {
	length      uint32
	ptype       uint8
	X           float64 `struc:"float64"`
	Y           float64 `struc:"float64"`
	Z           float64 `struc:"float64"`
	RX          float64 `struc:"float64"`
	RY          float64 `struc:"float64"`
	RZ          float64 `struc:"float64"`
	TCPOffsetX  float64 `struc:"float64"`
	TCPOffsetY  float64 `struc:"float64"`
	TCPOffsetZ  float64 `struc:"float64"`
	TCPOffsetRX float64 `struc:"float64"`
	TCPOffsetRY float64 `struc:"float64"`
	TCPOffsetRZ float64 `struc:"float64"`
}

type ConfigurationData struct { // 445 Bytes
	length             uint32
	ptype              uint8
	JointMinLimit      [6]float64 `struc:"[6]float64"`
	JointMaxLimit      [6]float64 `struc:"[6]float64"`
	JointMaxSpeed      [6]float64 `struc:"[6]float64"`
	JointMaxAcc        [6]float64 `struc:"[6]float64"`
	VJointDefault      float64    `struc:"float64"`
	AJointDefault      float64    `struc:"float64"`
	VToolDefault       float64    `struc:"float64"`
	AToolDefault       float64    `struc:"float64"`
	EqRadius           float64    `struc:"float64"`
	DHa                [6]float64 `struc:"[6]float64"`
	DHd                [6]float64 `struc:"[6]float64"`
	DHalpha            [6]float64 `struc:"[6]float64"`
	DHtheta            [6]float64 `struc:"[6]float64"`
	MasterboardVersion int32      `struc:"int32"`
	ControllerBoxType  int32      `struc:"int32"`
	RobotType          int32      `struc:"int32"`
	RobotSubType       int32      `struc:"int32"`
}

type ForceModeData struct { // 61 Bytes
	length    uint32
	ptype     uint8
	X         float64 `struc:"float64"`
	Y         float64 `struc:"float64"`
	Z         float64 `struc:"float64"`
	RX        float64 `struc:"float64"`
	RY        float64 `struc:"float64"`
	RZ        float64 `struc:"float64"`
	Dexterity float64 `struc:"float64"`
}

type AdditionalInfo struct {
	length                 uint32
	ptype                  uint8
	FreedriveButtonPressed bool `struc:"bool"`
	FreedriveButtonEnabled bool `struc:"bool"`
	IOEnabledFreedrive     bool `struc:"bool"`
}
