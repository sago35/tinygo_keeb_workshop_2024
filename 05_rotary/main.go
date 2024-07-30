package main

import (
	"machine"

	"tinygo.org/x/drivers/encoders"
)

func main() {
	enc := encoders.NewQuadratureViaInterrupt(
		machine.GPIO3,
		machine.GPIO4,
	)

	enc.Configure(encoders.QuadratureConfig{
		Precision: 4,
	})

	for oldValue := 0; ; {
		if newValue := enc.Position(); newValue != oldValue {
			println("value: ", newValue)
			oldValue = newValue
		}
	}

}
