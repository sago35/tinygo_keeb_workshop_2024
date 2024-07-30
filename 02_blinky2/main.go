package main

import (
	"image/color"
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

func (ws *WS2812B) PutColor(c color.Color) {
	ws.ws.PutColor(c)
}

func (ws *WS2812B) WriteRaw(rawGRB []uint32) error {
	return ws.ws.WriteRaw(rawGRB)
}

func main() {
	ws := NewWS2812B(machine.GPIO1)

	colors := [2][]uint32{
		{
			0xFFFFFFFF, 0xFF0000FF, 0x00FF00FF, 0x0000FFFF,
			0xFFFFFFFF, 0xFF0000FF, 0x00FF00FF, 0x0000FFFF,
			0xFFFFFFFF, 0xFF0000FF, 0x00FF00FF, 0x0000FFFF,
		},
		{
			0x00000000, 0x00000000, 0x00000000, 0x00000000,
			0x00000000, 0x00000000, 0x00000000, 0x00000000,
			0x00000000, 0x00000000, 0x00000000, 0x00000000,
		},
	}
	for {
		for i := range colors[0] {
			time.Sleep(time.Millisecond * 100)
			ws.WriteRaw(colors[0][:i+1])
		}
		time.Sleep(time.Millisecond * 500)
		ws.WriteRaw(colors[1])
	}
}
