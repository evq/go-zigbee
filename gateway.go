package zigbee

import (
	"time"
}

type ZigbeeMsg interface {
	ToBytes() []byte
}

type JoinResponse struct {
	NetAddr       uint16
	IeeeAddr      [8]uint8
}

type NodeDescriptorResponse struct {
	ManufacturerCode uint16
}

type ActiveEndpointResponse struct {
	Endpoints []Endpoint
}

type SimpleDescriptorResponse struct {
	Endpoint
}

type BindResponse struct {
	Success bool
}

type AttributeResponse struct {
  ID uint16
	Value string
}

type DeviceGateway interface {
  Connect(address string) error
	SendAsync(msg ZigbeeMsg) error
	Send(msg ZigbeeMsg) error
	PermitJoining(duration time.Duration) (ZigbeeMsg, chan JoinResponse, error)
	NodeDescriptorReq(z ZigbeeDevice) (ZigbeeMsg, chan NodeDescriptorResponse, error)
	ActiveEndpointReq(z ZigbeeDevice) (ZigbeeMsg, chan ActiveEndpointResponse, error)
	SimpleDescriptorReq(z ZigbeeDevice) (ZigbeeMsg, chan SimpleDescriptorResponse, error)
	BindReq(z ZigbeeDevice, e Endpoint, c Cluster) (ZigbeeMsg, chan BindResponse, error)
	ReadAttributes(z ZigbeeDevice, e Endpoint, c Cluster, attribids []uint8) (ZigbeeMsg, chan AttributeResponse, error)
}

type HAGateway interface {
  Connect(address string) error
	SendAsync(msg ZigbeeMsg) (chan ZigbeeMsg, error)
	Send(msg ZigbeeMsg) (ZigbeeMsg, error)
	SetOnOff(z ZigbeeDevice, endpointid uint8, value uint8) (ZigbeeMsg, error)
	SetLevel(z ZigbeeDevice, endpointid uint8, value uint8, transitiontime uint16) (ZigbeeMsg, error)
}
