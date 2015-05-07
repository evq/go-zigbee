package zigbee

import (
	"time"
)

type JoinResponse struct {
	NetAddr  uint16
	IeeeAddr [8]uint8
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
	ID    uint16
	Value string
}

type DeviceGateway interface {
	Connect(address string) error
	SendAsync() error
	Send() error
	PermitJoining(duration time.Duration) (chan JoinResponse, error)
	NodeDescriptorReq(z ZigbeeDevice) (chan NodeDescriptorResponse, error)
	ActiveEndpointReq(z ZigbeeDevice) (chan ActiveEndpointResponse, error)
	SimpleDescriptorReq(z ZigbeeDevice) (chan SimpleDescriptorResponse, error)
	BindReq(z ZigbeeDevice, e Endpoint, c Cluster) (chan BindResponse, error)
	ReadAttributes(z ZigbeeDevice, e Endpoint, c Cluster, attribids []uint8) (chan AttributeResponse, error)
}

type HAGateway interface {
	Connect(address string) error
	SendAsync() error
	Send() error
	SetOnOff(z ZigbeeDevice, endpointid uint8, value uint8) error
	SetLightLevel(z ZigbeeDevice, endpointid uint8, level uint8, transitiontime uint16) error
}
