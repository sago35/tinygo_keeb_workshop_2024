package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	ax := machine.ADC{Pin: machine.GPIO29}
	ax.Configure(machine.ADCConfig{})
	ay := machine.ADC{Pin: machine.GPIO28}
	ay.Configure(machine.ADCConfig{})

	btn := machine.GPIO0
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		x := ax.Get()
		y := ay.Get()
		pressed := !btn.Get()
		fmt.Printf("%04X %04X %t\n", x, y, pressed)
		time.Sleep(200 * time.Millisecond)
	}
}
