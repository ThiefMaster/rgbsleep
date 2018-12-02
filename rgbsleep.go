package main

import (
	"log"

	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	log.Println("initializing")
	leds.Init()

	lockedChan := wts.RunMonitor()
	for {
		if <-lockedChan {
			log.Println("session locked")
			leds.TurnOff()
		} else {
			log.Println("session unlocked")
			leds.TurnOn()
		}
	}
}
