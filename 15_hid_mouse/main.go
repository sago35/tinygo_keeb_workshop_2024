package main

import (
	"machine"
	"machine/usb/hid/mouse"
)

func main() {
	btn := machine.GPIO2
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	m := mouse.Port()

	for {
		if !btn.Get() {
			m.Press(mouse.Left)
		} else {
			m.Release(mouse.Left)
		}
	}
}
