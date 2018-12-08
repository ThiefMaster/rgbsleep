package leds

var fadeAuraControl = make(chan bool)

func Init() {
	InitAura()
	InitMystic()
	// this program runs after login, so ensure the LEDs are on.
	TurnAuraOn()
	TurnMysticOn()
	go RunAuraFader(fadeAuraControl)
}

func TurnOn() {
	TurnMysticOn()
	fadeAuraControl <- true
}

func TurnOff() {
	TurnMysticOff()
	fadeAuraControl <- false
}
