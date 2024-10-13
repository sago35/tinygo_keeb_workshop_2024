package main

import (
	"fmt"
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

type WS2812B struct {
	Pin machine.Pin
	ws  *piolib.WS2812B
}

func NewWS2812B(pin machine.Pin) *WS2812B {
	s, _ := pio.PIO0.ClaimStateMachine()
	ws, _ := piolib.NewWS2812B(s, pin)
	ws.EnableDMA(true)
	return &WS2812B{
		ws: ws,
	}
}

func (ws *WS2812B) WriteRaw(rawGRB []uint32) error {
	return ws.ws.WriteRaw(rawGRB)
}

func main() {
	colors := []uint32{
		0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
		0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
		0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
	}

	ws := NewWS2812B(machine.GPIO1)

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
		ws.WriteRaw(colors)

		// COL1
		colPins[0].High()
		colPins[1].Low()
		colPins[2].Low()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw1 pressed\n")
			colors[0] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[0] = 0xFFFFFFFF
		}
		if rowPins[1].Get() {
			fmt.Printf("sw5 pressed\n")
			colors[1] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[1] = 0xFFFFFFFF
		}
		if rowPins[2].Get() {
			fmt.Printf("sw9 pressed\n")
			colors[2] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[2] = 0xFFFFFFFF
		}

		// COL2
		colPins[0].Low()
		colPins[1].High()
		colPins[2].Low()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw2 pressed\n")
			colors[3] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[3] = 0xFFFFFFFF
		}
		if rowPins[1].Get() {
			fmt.Printf("sw6 pressed\n")
			colors[4] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[4] = 0xFFFFFFFF
		}
		if rowPins[2].Get() {
			fmt.Printf("sw10 pressed\n")
			colors[5] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[5] = 0xFFFFFFFF
		}

		// COL3
		colPins[0].Low()
		colPins[1].Low()
		colPins[2].High()
		colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw3 pressed\n")
			colors[6] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[6] = 0xFFFFFFFF
		}
		if rowPins[1].Get() {
			fmt.Printf("sw7 pressed\n")
			colors[7] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[7] = 0xFFFFFFFF
		}
		if rowPins[2].Get() {
			fmt.Printf("sw11 pressed\n")
			colors[8] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[8] = 0xFFFFFFFF
		}

		// COL4
		colPins[0].Low()
		colPins[1].Low()
		colPins[2].Low()
		colPins[3].High()
		time.Sleep(1 * time.Millisecond)

		if rowPins[0].Get() {
			fmt.Printf("sw4 pressed\n")
			colors[9] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[9] = 0xFFFFFFFF
		}
		if rowPins[1].Get() {
			fmt.Printf("sw8 pressed\n")
			colors[10] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[10] = 0xFFFFFFFF
		}
		if rowPins[2].Get() {
			fmt.Printf("sw12 pressed\n")
			colors[11] = 0x00000000
			ws.WriteRaw(colors)
			time.Sleep(100 * time.Millisecond)
		} else {
			colors[11] = 0xFFFFFFFF
		}
	}
}
