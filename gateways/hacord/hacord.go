package hacord

import (
	"net"
)

type HACordGateway struct {
	net.Conn
}

func (g HACordGateway)Connect(address string) error {
	g.Conn, err := net.Dial("tcp", address)
	return err
}

func SendAsync(msg ZigbeeMsg) (chan ZigbeeMsg, error) {
	// Not implemented
	return nil, nil
}


