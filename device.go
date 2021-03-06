package zigbee

type ZigbeeDevice struct {
	Name             string   `json:"name"`
	Endpoints        map[uint8]*Endpoint `json:"-"`
	NetAddr          uint16  `json:"netaddr"`
	IeeeAddr         uint64 `json:"ieeeaddr"`
	ManufacturerCode uint16 `json:"-"`
}

// ZLL DeviceTypes
const (
	ZLLOnOffLight = 0x0000
	ZLLOnOffPlug = 0x0010
	ZLLDimmableLight = 0x0100
	HAOnOffLight = 0x0100
	HADimmableLight = 0x0101
	HAColorLight = 0x0102
	ZLLDimmablePlug = 0x0110
	ZLLColorLight = 0x0200
	ZLLExtendedColorLight = 0x0210
	ZLLColorTemperatureLight = 0x0220
)

// ProfileIds
const (
	HomeAutomation = 0x0104
	LightLink      = 0xC05E
)

type Endpoint struct {
	ID          uint8
	Profile     uint16
	DeviceType  uint16
	InClusters  map[uint16]*Cluster
	OutClusters map[uint16]*Cluster
}

// ZCL Clusters
const (
	BasicCluster     = 0x0000
	IdentifyCluster  = 0x0003
	GroupsCluster    = 0x0004
	ScenesCluster    = 0x0005
	OnOffCluster     = 0x0006
	LevelCluster     = 0x0008
)

type Cluster struct {
	ID         uint16
	Attributes map[uint16]string
}

// BasicAttributes
const (
	ZCLVersion         = iota // 0
	ApplicationVersion        // 1 ...
	StackVersion
	HWVersion
	ManufacturerName
	ModelId
	DateCode
	PowerSource
)

// GroupAttributes
const NameSupport = 0

// On/Off Attributes
const OnOff = 0

//OnOff Values
const (
	Off = 0x00
	On  = 0x01
)

// Level Attributes
const Level = 0
