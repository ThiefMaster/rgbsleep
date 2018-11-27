package main

import (
	"log"

	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	log.Println("initializing")
	leds.Init()

	for {
		select {
		case locked := <-wts.RunMonitor():
			if locked {
				log.Println("session locked")
				leds.TurnOff()
			} else {
				log.Println("session unlocked")
				leds.TurnOn()
			}
		}
	}
}
