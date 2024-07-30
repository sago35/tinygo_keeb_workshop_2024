package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
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

	invDisplay := InvertedDisplay{&display}

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}

	tinyfont.WriteLine(&invDisplay, &freemono.Bold9pt7b, 5, 10, "hello", white)
	tinyfont.WriteLine(&invDisplay, &gophers.Regular32pt, 5, 50, "ABCEF", white)
	display.Display()
}

type InvertedDisplay struct {
	drivers.Displayer
}

func (d *InvertedDisplay) SetPixel(x, y int16, c color.RGBA) {
	sx, sy := d.Displayer.Size()
	d.Displayer.SetPixel(sx-x, sy-y, c)
}
