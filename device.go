package zigbee

type ZigbeeDevice struct {
	Name             string   `json:"name"`
	Endpoints        []Endpoint `json:"-"`
	NetAddr          uint16  `json:"netaddr"`
	IeeeAddr         [8]uint8 `json:"ieeeaddr"`
	ManufacturerCode uint16 `json:"-"`
}

// ZLL DeviceTypes
const (
	OnOffLight = 0x0000
	OnOffPlug = 0x0010
	DimmableLight = 0x0100
	DimmablePlug = 0x0110
	ColorLight = 0x0200
	ExtendedColorLight = 0x0210
	ColorTemperatureLight = 0x0220
)

// HA DeviceTypes
//const (
	//OnOffLight = 0x0100
	//DimmableLight = 0x0101
	//ColorLight = 0x0102
//)

// ProfileIDs
const (
	HomeAutomation = 0x0104
	LightLink      = 0xC05E
)

type Endpoint struct {
	ID          uint8
	Profile     uint16
	DeviceType  uint16
	InClusters  []Cluster
	OutClusters []Cluster
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
	ModelID
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
