package main

import (
	"fmt"
	"github.com/evq/go-zigbee"
	"github.com/evq/go-zigbee/gateways/embercli"
	"time"
)

func main() {
	addr := "192.168.2.177:4901"

	var z zigbee.HAGateway = &embercli.EmberCliGateway{}

	fmt.Println("Connecting to " + addr)
	err := z.Connect(addr)

	if err != nil {
		fmt.Println("Error connecting!")
	}
	fmt.Println("Connected")

	dev := zigbee.ZigbeeDevice{}
	dev.NetAddr = 0x1a81
	dev.IeeeAddr = [8]byte{0x00, 0x17, 0x88, 0x01, 0x00, 0xfc, 0x42, 0x88}
	var endpoint uint8 = 0x0b

	err = z.SetOnOff(dev, endpoint, zigbee.On)
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

	time.Sleep(2 * time.Second)

	err = z.MoveToHueSat(dev, endpoint, 0x00, 0xff, 10)
	if err != nil {
		fmt.Println("Error serializing payload!")
		fmt.Println(err)
	}

	fmt.Println("Transitioning light to red")
	err = z.Send()
	if err != nil {
		fmt.Println("Error sending!")
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	err = z.MoveToHueSat(dev, endpoint, 0x54, 0xff, 10)
	if err != nil {
		fmt.Println("Error serializing payload!")
		fmt.Println(err)
	}

	fmt.Println("Transitioning light to green")
	err = z.Send()
	if err != nil {
		fmt.Println("Error sending!")
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	err = z.MoveToHueSat(dev, endpoint, 0xA8, 0xff, 10)
	if err != nil {
		fmt.Println("Error serializing payload!")
		fmt.Println(err)
	}

	fmt.Println("Transitioning light to blue")
	err = z.Send()
	if err != nil {
		fmt.Println("Error sending!")
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	err = z.SetOnOff(dev, endpoint, zigbee.Off)
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
