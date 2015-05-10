package embercli

import (
	"fmt"
	"github.com/evq/go-zigbee"
	"net"
)

type EmberCliGateway struct {
	Conn       net.Conn `json:"-"`
	TXBuffer   EmberCliCmd `json:"-"`
	Address    string `json:"address"`
}

type EmberCliCmd struct {
	Cmd string
	NetAddr          uint16
	Endpoint uint8
}


func (g *EmberCliGateway) Connect(address string) error {
	var err error
	g.Conn, err = net.Dial("tcp", address)
	return err
}

func (g *EmberCliGateway) Reconnect() error {
	var err error
	g.Conn, err = net.Dial("tcp", g.Address)
	return err
}

func (g *EmberCliGateway) SendAsync() error {
	// Not implemented
	return nil
}

func (g *EmberCliGateway) Send() error {

	//fmt.Printf(g.TXBuffer.Cmd)
	_, err := g.Conn.Write([]byte(g.TXBuffer.Cmd))
	if err != nil {
		return err
	}

	//fmt.Printf("send 0x%x 0x%x 0x%x\n", g.TXBuffer.NetAddr, 0x01, g.TXBuffer.Endpoint)
	_, err = g.Conn.Write([]byte(fmt.Sprintf("send 0x%x 0x%x 0x%x\n", g.TXBuffer.NetAddr, 0x01, g.TXBuffer.Endpoint)))

	return err
}

func (g *EmberCliGateway) SetOnOff(z zigbee.ZigbeeDevice, endpointid uint8, value uint8) error {
	strval := "off"
	if value == zigbee.On {
		strval = "on"
	}

	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl on-off %s\n", strval)

	return nil
}

func (g *EmberCliGateway) MoveToLightLevelWOnOff(z zigbee.ZigbeeDevice, endpointid uint8, level uint8, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl level-control o-mv-to-level 0x%x 0x%x\n", level, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToLightLevel(z zigbee.ZigbeeDevice, endpointid uint8, level uint8, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl level-control mv-to-level 0x%x 0x%x\n", level, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToHue(z zigbee.ZigbeeDevice, endpointid uint8, hue uint8, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol movetohue 0x%x 0x%x\n", hue, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToSat(z zigbee.ZigbeeDevice, endpointid uint8, sat uint8, transitiontime uint16) error {
	if sat == 0xff {
		sat = 0xfe
	}
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol movetosat 0x%x 0x%x\n", sat, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToHueSat(z zigbee.ZigbeeDevice, endpointid uint8, hue uint8, sat uint8, transitiontime uint16) error {
	if sat == 0xff {
		sat = 0xfe
	}
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol movetohueandsat 0x%x 0x%x 0x%x\n", hue, sat, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToXY(z zigbee.ZigbeeDevice, endpointid uint8, X uint16, Y uint16, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol movetocolor 0x%x 0x%x 0x%x\n", X, Y, transitiontime)

	return nil
}

func (g *EmberCliGateway) MoveToColorTemp(z zigbee.ZigbeeDevice, endpointid uint8, temp uint16, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol movetocolortemp 0x%x 0x%x\n", temp, transitiontime)

	return nil
}

func (g *EmberCliGateway) Loop(z zigbee.ZigbeeDevice, endpointid uint8, starthue uint16, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol loop 0xff 0x01 0x01 0x%x 0x%x\n", transitiontime, starthue)

	return nil
}

func (g *EmberCliGateway) StopLoop(z zigbee.ZigbeeDevice, endpointid uint8, starthue uint16, transitiontime uint16) error {
	g.TXBuffer.NetAddr = z.NetAddr
	g.TXBuffer.Endpoint = endpointid
	g.TXBuffer.Cmd = fmt.Sprintf("zcl colorcontrol loop 0xff 0x00 0x01 0x%x 0x%x\n", transitiontime, starthue)

	return nil
}
