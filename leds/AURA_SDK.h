#include <Windows.h>

typedef void* MbLightControl;

typedef DWORD (WINAPI* EnumerateMbControllerFunc)(MbLightControl handles[], DWORD size);
typedef DWORD (WINAPI* SetMbModeFunc) (MbLightControl handle, DWORD mode);
typedef DWORD (WINAPI* SetMbColorFunc) (MbLightControl handle, BYTE* color, DWORD size);
typedef DWORD (WINAPI* GetMbLedCountFunc)(MbLightControl handle);
