package hacord

//Cmds
const (
	OnOff      = 0x0400
	LightLevel = 0x0c00
)

//OnOff Values
const (
	HACordOff = 0x01
	HACordOn  = 0x00
)

type OnOffPacket struct {
	HACordHeader
	MsgLen  int `struc:"uint8,sizeof=Payload"`
	Payload OnOffPayload
}

type OnOffPayload struct {
	State        uint8    `struc:uint8`
	NetAddr      uint16   `struc:uint16`
	IeeeAddr     [8]uint8 `struc:[8]uint8`
	SrcEndpoint  uint8    `struc:uint8`
	DestEndpoint uint8    `struc:uint8`
	GroupId      uint16   `struc:uint16`
}

type LightLevelPacket struct {
	HACordHeader
	MsgLen  int `struc:"uint8,sizeof=Payload"`
	Payload LightLevelPayload
}

type LightLevelPayload struct {
	Level          uint8    `struc:uint8`
	TransitionTime uint16   `struc:uint16`
	NetAddr        uint16   `struc:uint16`
	IeeeAddr       [8]uint8 `struc:[8]uint8`
	SrcEndpoint    uint8    `struc:uint8`
	DestEndpoint   uint8    `struc:uint8`
	GroupId        uint16   `struc:uint16`
}

type PermitJoiningPayload struct {
	DurationInSeconds uint8
	Magic             uint8
}

type RequestNodeDescriptionPayload struct {
	NetAddr uint16
}

type ReadAttributesPayload struct {
	NetAddr   uint16
	LongAddr  [8]uint8
	Enpt      uint8
	ClusterID uint16
	AttrLen   uint8
	Attr      []uint16
}

type RequestActiveEndpointPayload struct {
	NetAddr uint16
}

type RequestSimpleDescriptionPayload struct {
	NetAddr  uint16
	Endpoint uint8
}

type RequestZDOBind struct {
	NetAddr   uint16
	LongAddr  [8]uint8
	Endpoint  uint8
	ClusterID uint16
	Magic     [9]uint8
}
