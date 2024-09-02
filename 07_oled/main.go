package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 2.8 * machine.MHz,
		SDA:       machine.GPIO12,
		SCL:       machine.GPIO13,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	black := color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}

	cnt := 0
	for {
		c := white
		if cnt == 1 {
			c = black
		}
		for x := int16(0); x < 128; x += 2 {
			for y := int16(0); y < 64; y += 2 {
				display.SetPixel(x+0, y+0, c)
				display.SetPixel(x+0, y+1, c)
				display.SetPixel(x+1, y+0, c)
				display.SetPixel(x+1, y+1, c)
				display.Display()
			}
		}
		cnt++
	}
}
