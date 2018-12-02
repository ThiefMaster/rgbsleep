#include <Windows.h>

typedef void* MbLightControl;
typedef void* GPULightControl;
typedef void* ClaymoreKeyboardLightControl;
typedef void* RogMouseLightControl;

DWORD EnumerateMbController(MbLightControl handles[], DWORD size);
DWORD SetMbMode(MbLightControl handle, DWORD mode);
DWORD SetMbColor(MbLightControl handle, BYTE* color, DWORD size);
DWORD GetMbColor(MbLightControl handle, BYTE* color, DWORD size);
DWORD GetMbLedCount(MbLightControl handle);

DWORD EnumerateGPU(GPULightControl handles[], DWORD size);
DWORD SetGPUMode(GPULightControl handle, DWORD mode);
DWORD SetGPUColor(GPULightControl handle, BYTE* color, DWORD size);
DWORD GetGPULedCount(GPULightControl handle);

DWORD CreateClaymoreKeyboard(ClaymoreKeyboardLightControl* handle);
DWORD SetClaymoreKeyboardMode(ClaymoreKeyboardLightControl handle, DWORD mode);
DWORD SetClaymoreKeyboardColor(ClaymoreKeyboardLightControl handle, BYTE* color, DWORD size);
DWORD GetClaymoreKeyboardLedCount(ClaymoreKeyboardLightControl handle);

DWORD CreateRogMouse(RogMouseLightControl* handle);
DWORD SetRogMouseMode(RogMouseLightControl handle, DWORD mode);
DWORD SetRogMouseColor(RogMouseLightControl handle, BYTE* color, DWORD size);
DWORD RogMouseLedCount(RogMouseLightControl handle);
