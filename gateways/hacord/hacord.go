package hacord

import (
	"bytes"
	"fmt"
	"github.com/evq/go-zigbee"
	"github.com/evq/struc"
	"net"
)

// MsgTypes
const (
	Command      = 0x43
	HACordAck    = 0x52
	NodeResponse = 0x42
)

type HACordHeader struct {
	MsgType  uint8  `struc:uint8`
	Sequence uint8  `struc:uint8`
	Cmd      uint16 `struc:uint16`
}

type HACordGateway struct {
	Conn       net.Conn
	CurrentCmd uint16
	Sequence   uint8
	TXBuffer   bytes.Buffer
}

func (g *HACordGateway) Connect(address string) error {
	var err error
	g.Conn, err = net.Dial("tcp", address)
	return err
}

func (g *HACordGateway) SendAsync() error {
	// Not implemented
	return nil
}

func (g *HACordGateway) Send() error {
	fmt.Println(g.TXBuffer.Bytes())

	_, err := g.Conn.Write(g.TXBuffer.Bytes())

	g.Sequence++
	g.TXBuffer.Reset()

	return err
}

func (g *HACordGateway) SetOnOff(z zigbee.ZigbeeDevice, endpointid uint8, value uint8) error {
	if value == zigbee.On {
		value = HACordOn
	} else {
		value = HACordOff
	}

	payload := OnOffPacket{
		HACordHeader{Command, g.Sequence, OnOff}, 0,
		OnOffPayload{value, z.NetAddr, z.IeeeAddr, 0x01, endpointid, 0x0000},
	}

	err := struc.Pack(&g.TXBuffer, &payload)
	return err
}

func (g *HACordGateway) SetLightLevel(z zigbee.ZigbeeDevice, endpointid uint8, level uint8, transitiontime uint16) error {
	payload := LightLevelPacket{
		HACordHeader{Command, g.Sequence, LightLevel}, 0,
		LightLevelPayload{level, transitiontime, z.NetAddr, z.IeeeAddr, 0x01, endpointid, 0x0000},
	}

	err := struc.Pack(&g.TXBuffer, &payload)
	return err
}
