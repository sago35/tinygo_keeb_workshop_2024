package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
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

	data := []byte("ABCEF")
	for {
		display.ClearBuffer()
		data[0], data[1], data[2], data[3], data[4] = data[1], data[2], data[3], data[4], data[0]
		tinyfont.WriteLine(&invDisplay, &gophers.Regular32pt, 5, 45, string(data), white)
		display.Display()
		time.Sleep(200 * time.Millisecond)
	}
}

type InvertedDisplay struct {
	drivers.Displayer
}

func (d *InvertedDisplay) SetPixel(x, y int16, c color.RGBA) {
	sx, sy := d.Displayer.Size()
	d.Displayer.SetPixel(sx-x, sy-y, c)
}
