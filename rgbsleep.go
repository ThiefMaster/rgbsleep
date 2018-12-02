package main

import (
	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	leds.Init()

	lockedChan := wts.RunMonitor()
	for {
		if <-lockedChan {
			leds.TurnOff()
		} else {
			leds.TurnOn()
		}
	}
}
