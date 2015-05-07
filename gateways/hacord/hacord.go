package hacord

import (
	"net"
)

type HACordGateway struct {
	net.Conn
}

func (g HACordGateway) Connect(address string) error {
	g.Conn, err := net.Dial("tcp", address)
	return err
}

func SendAsync() error {
	// Not implemented
	return nil, nil
}


