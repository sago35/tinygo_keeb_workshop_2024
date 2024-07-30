package main

import (
	"machine"
	"time"
)

func main() {
	btn := machine.GPIO2
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if !btn.Get() {
			println("pressed")
			time.Sleep(10 * time.Millisecond)

			for !btn.Get() {
			}
			println("released")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
