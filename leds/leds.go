package leds

/*
#cgo windows LDFLAGS: MysticLight_SDK.dll

#include "MysticLight_SDK.h"

void TurnOn() {
	MLAPI_SetLedStyle(L"MSI_VGA", 0, L"No animation");
	MLAPI_SetLedColor(L"MSI_VGA", 0, 255, 50, 0);
	MLAPI_SetLedBright(L"MSI_VGA", 0, 5);
}

void TurnOff() {
	MLAPI_SetLedStyle(L"MSI_VGA", 0, L"Off");
}
*/
import "C"

var fadeState = make(chan bool)

func Init() {
	go RunAuraFader(fadeState)
	C.MLAPI_Initialize()
}

func TurnOn() {
	C.TurnOn()
	fadeState <- true
}

func TurnOff() {
	C.TurnOff()
	fadeState <- false
}
