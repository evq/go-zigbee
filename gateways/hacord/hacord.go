package hacord

import (
	"github.com/evq/go-zigbee"
	"bytes"
	"net"
	"encoding/binary"
)

// MsgTypes
const (
	Command = 0x43
	HACordAck = 0x52
	NodeResponse = 0x42
)

//Cmds
const (
	OnOff = 0x0400
)

//OnOff Values
const (
	HACordOff = 0x01
	HACordOn = 0x00
)

type HACordPacket struct {
	MsgType uint8
	Sequence uint8
	Cmd uint16
	MsgLen uint8
}

type OnOffPayload struct {
	Val uint8
	NetAddr       uint16
	IeeeAddr      [8]uint8
	SrcEndpoint uint8
	DestEndpoint uint8
	GroupId uint16
}

type HACordGateway struct {
	Conn net.Conn
	CurrentCmd uint16
	Sequence uint8
	TXBuffer []byte
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
	pkt := HACordPacket{Command, g.Sequence, g.CurrentCmd, uint8(len(g.TXBuffer))}

	err := binary.Write(g.Conn, binary.BigEndian, pkt)
	if err != nil {
		return err
	}
	_, err = g.Conn.Write(g.TXBuffer)
	if err != nil {
		return err
	}
	g.Sequence++
	return nil
}

func (g *HACordGateway) SetOnOff(z zigbee.ZigbeeDevice, endpointid uint8, value uint8) error {
	if value == zigbee.On {
		value = HACordOn
	} else {
		value = HACordOff
	}

	payload := OnOffPayload{value, z.NetAddr, z.IeeeAddr, 0x01, endpointid, 0x0000}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, payload)
	if err != nil {
		return err
	}
	g.TXBuffer = buf.Bytes()
	g.CurrentCmd = OnOff
	return nil
}

func (g *HACordGateway) SetLevel(z zigbee.ZigbeeDevice, endpointid uint8, value uint8, transitiontime uint16) error {
	// Not implemented
	return nil
}
