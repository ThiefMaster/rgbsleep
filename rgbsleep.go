package main

import (
	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	leds.Init()

	monitor := wts.RunMonitor()
	for {
		select {
		case locked := <-monitor:
			if locked {
				leds.TurnOff()
			} else {
				leds.TurnOn()
			}
		}
	}
}
