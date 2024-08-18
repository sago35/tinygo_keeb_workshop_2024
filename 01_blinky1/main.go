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
	return &WS2812B{
		ws: ws,
	}
}

func (ws *WS2812B) PutColor(c color.Color) {
	ws.ws.PutColor(c)
}

var (
	white = color.RGBA{R: 0x20, G: 0x20, B: 0x20, A: 0x00}
	black = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
)

func main() {
	ws := NewWS2812B(machine.GPIO16)
	ws.PutColor(white)

	for {
		time.Sleep(time.Millisecond * 500)
		ws.PutColor(black)
		time.Sleep(time.Millisecond * 500)
		ws.PutColor(white)
	}
}
