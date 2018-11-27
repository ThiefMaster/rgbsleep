package main

import (
	"log"

	"github.com/thiefmaster/rgbsleep/leds"
	"github.com/thiefmaster/rgbsleep/wts"
)

func main() {
	leds.Init()

	changes := make(chan wts.Message, 100)
	closeChan := make(chan int)

	wts.Subscribe(changes, closeChan)

	for {
		select {
		case msg := <-changes:
			switch msg.Param {
			case wts.WTS_SESSION_LOCK:
				log.Println("session locked")
				leds.TurnOff()
			case wts.WTS_SESSION_UNLOCK:
				log.Println("session unlocked")
				leds.TurnOn()
			}
			close(msg.ChanOk)
		}
	}
}
