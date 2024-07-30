package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	colPins := []machine.Pin{
		machine.GPIO5,
		machine.GPIO6,
		machine.GPIO7,
		machine.GPIO8,
	}

	rowPins := []machine.Pin{
		machine.GPIO9,
		machine.GPIO10,
		machine.GPIO11,
	}

	for _, c := range colPins {
		c.Configure(machine.PinConfig{Mode: machine.PinOutput})
		c.Low()
	}

	for _, c := range rowPins {
		c.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}

	for {
		// COL1
		colPins[0].High()
		colPins[1].Low()
		colPins[2].Low()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw1 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[1].Get() {
			fmt.Printf("sw5 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[2].Get() {
			fmt.Printf("sw9 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}

		// COL2
		colPins[0].Low()
		colPins[1].High()
		colPins[2].Low()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw2 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[1].Get() {
			fmt.Printf("sw6 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[2].Get() {
			fmt.Printf("sw10 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}

		// COL3
		colPins[0].Low()
		colPins[1].Low()
		colPins[2].High()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw3 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[1].Get() {
			fmt.Printf("sw7 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[2].Get() {
			fmt.Printf("sw11 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}

		// COL4
		colPins[0].Low()
		colPins[1].Low()
		colPins[2].Low()
		colPins[3].High()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw4 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[1].Get() {
			fmt.Printf("sw8 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
		if rowPins[2].Get() {
			fmt.Printf("sw12 pressed\n")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
