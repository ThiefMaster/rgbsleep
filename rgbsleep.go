package main

import (
	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	leds.Init()

	for {
		select {
		case locked := <-wts.RunMonitor():
			if locked {
				leds.TurnOff()
			} else {
				leds.TurnOn()
			}
		}
	}
}
