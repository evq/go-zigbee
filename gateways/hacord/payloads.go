package hacord

import (
  "fmt"
  "net"
  "log"
  "time"
)

// Message type

type PermitJoiningPayload struct {
  DurationInSeconds uint8
  Magic uint8
}

type RequestNodeDescriptionPayload struct {
  NetAddr uint16
}

type ReadAttributesPayload struct {
  NetAddr uint16
  LongAddr [8]uint8
  Enpt uint8
  ClusterID uint16
  AttrLen uint8
  Attr []uint16
}

type RequestActiveEndpointPayload struct {
  NetAddr uint16
}

type RequestSimpleDescriptionPayload struct {
  NetAddr uint16
  Endpoint uint8
}

type RequestZDOBind struct {
  NetAddr uint16
  LongAddr [8]uint8
  Endpoint uint8
  ClusterID uint16
  Magic [9]uint8
}
