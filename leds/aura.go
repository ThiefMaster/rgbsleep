package leds

/*
#include "AURA_SDK.h"

static EnumerateMbControllerFunc EnumerateMbController;
static SetMbModeFunc SetMbMode;
static SetMbColorFunc SetMbColor;
static GetMbLedCountFunc GetMbLedCount;

static MbLightControl handle;
static DWORD ledCount;
static BYTE *colors;
static BYTE ready;


static void initDLL() {
	HMODULE dll = LoadLibrary("AURA_SDK.dll");
	EnumerateMbControllerFunc EnumerateMbController = (EnumerateMbControllerFunc)GetProcAddress(dll, "EnumerateMbController");
	SetMbModeFunc SetMbMode = (SetMbModeFunc)GetProcAddress(dll, "SetMbMode");
	SetMbColorFunc SetMbColor = (SetMbColorFunc)GetProcAddress(dll, "SetMbColor");
	GetMbLedCountFunc GetMbLedCount = (GetMbLedCountFunc)GetProcAddress(dll, "GetMbLedCount");
}

BYTE Init() {
	initDLL();
	DWORD count = EnumerateMbController(NULL, 0);
	if (count != 1) {
		return 1;
	}
	return 2;
	MbLightControl *handles = calloc(count, sizeof(MbLightControl));
	EnumerateMbController(handles, count);
	handle = handles[0];
	SetMbMode(handle, 1);
	ledCount = GetMbLedCount(handle);
	colors = calloc(3 * ledCount, sizeof(BYTE));
	ready = 1;
	return 0;
}

void SetColors(BYTE boardR, BYTE boardG, BYTE boardB, BYTE caseR, BYTE caseG, BYTE caseB) {
	if (!ready) {
		return;
	}
	for (int i = 0; i < ledCount; i++) {
		BYTE r, g, b;
		if (i == ledCount - 1) {
			// rgb header for front panel controller
			r = caseR;
			g = caseG;
			b = caseB;
		} else if (i == ledCount - 2) {
			// unused rgb header
			r = 0;
			g = 0;
			b = 0;
		} else {
			// onboard led
			r = boardR;
			g = boardG;
			b = boardB;
		}
		colors[i*3 + 0] = r;
		colors[i*3 + 1] = g;
		colors[i*3 + 2] = b;
	}
	SetMbColor(handle, colors, ledCount * 3 * sizeof(BYTE));
}
*/
import "C"

import (
	"image/color"
	"log"
	"time"
)

var (
	colorOff     = color.RGBA{}
	colorOnCase  = color.RGBA{R: 255, G: 50, B: 0}
	colorOnBoard = color.RGBA{R: 255, G: 185, B: 15}
)

func fade(from, to color.RGBA, progress uint8, reverse bool) color.RGBA {
	if reverse {
		progress = 100 - progress
	}
	return color.RGBA{
		R: uint8(float32(from.R) + float32(progress)/100*float32(to.R-from.R)),
		G: uint8(float32(from.G) + float32(progress)/100*float32(to.G-from.G)),
		B: uint8(float32(from.B) + float32(progress)/100*float32(to.B-from.B)),
	}
}

func RunAuraFader(stateChan chan bool) {
	if res := int(C.Init()); res != 0 {
		log.Printf("aura init failed: %d\n", res)
	}
	for {
		var reverse bool
		if <-stateChan {
			log.Println("fade to on")
			reverse = false
		} else {
			log.Println("fade to off")
			reverse = true
		}
		for i := uint8(0); i <= 100; i++ {
			colorCase := fade(colorOff, colorOnCase, i, reverse)
			colorBoard := fade(colorOff, colorOnBoard, i, reverse)
			C.SetColors(
				C.uchar(colorBoard.R), C.uchar(colorBoard.G), C.uchar(colorBoard.B),
				C.uchar(colorCase.R), C.uchar(colorCase.G), C.uchar(colorCase.B),
			)
			log.Printf("step %d: %v | %v", i, colorCase, colorBoard)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
