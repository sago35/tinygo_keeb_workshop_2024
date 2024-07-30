package main

import (
	"machine"
	"machine/usb/hid/keyboard"
)

func main() {
	btn := machine.GPIO2
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := keyboard.Port()
	for {
		if !btn.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}
	}
}
