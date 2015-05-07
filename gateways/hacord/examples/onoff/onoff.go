package main

import (
	"github.com/evq/go-zigbee"
	"github.com/evq/go-zigbee/gateways/hacord"
	"fmt"
	"time"
)

func main() {
	addr := "192.168.2.177:1234"

	var z zigbee.HAGateway = &hacord.HACordGateway{}

	fmt.Println("Connecting to " + addr)
	err := z.Connect(addr)

	if err != nil {
		fmt.Println("Error connecting!")
	}
	fmt.Println("Connected")

  dev := zigbee.ZigbeeDevice{}
	dev.NetAddr = 0x98fd
	dev.IeeeAddr = [8]byte{0x7c,0xe5,0x24,0x00,0x00,0x01,0x02,0x37}

	err = z.SetOnOff(dev, 0x01, zigbee.On)
	if err != nil {
		fmt.Println("Error serializing payload!")
		fmt.Println(err)
	}
	fmt.Println("Turning light on")
	err = z.Send()
	if err != nil {
		fmt.Println("Error sending!")
		fmt.Println(err)
	}

	time.Sleep(2*time.Second)

	err = z.SetOnOff(dev, 0x01, zigbee.Off)
	if err != nil {
		fmt.Println("Error serializing payload!")
		fmt.Println(err)
	}
	fmt.Println("Turning light off")
	err = z.Send()
	if err != nil {
		fmt.Println("Error sending!")
		fmt.Println(err)
	}
}
