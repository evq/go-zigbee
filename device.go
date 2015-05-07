package zigbee

// ZCL Clusters
const (
	BasicCluster    uint16 = 0x0000
	IdentifyCluster uint16 = 0x0003
	GroupsCluster   uint16 = 0x0004
	ScenesCluster   uint16 = 0x0005
	OnOffCluster    uint16 = 0x0006
	LevelCluster    uint16 = 0x0008
)

type ZigbeeDevice struct {
	Name             string
	Endpoints        []Endpoint
	NetAddr          uint16
	IeeeAddr         [8]uint8
	ManufacturerCode uint16
}

type Endpoint struct {
	ID          uint8
	Profile     uint16
	DeviceType  uint16
	InClusters  []Cluster
	OutClusters []Cluster
}

// ProfileIDs
const (
	HomeAutomation = 0x0104
	LightLink      = 0xC05E
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
