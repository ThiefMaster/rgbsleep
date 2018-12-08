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

func InitMystic() {
	C.MLAPI_Initialize()
}

func TurnMysticOn() {
	C.TurnOn()
}

func TurnMysticOff() {
	C.TurnOff()
}
