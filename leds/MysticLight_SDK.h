#include <Windows.h>

int MLAPI_Initialize();
int MLAPI_GetErrorMessage(int ErrorCode, BSTR* pDesc);
int MLAPI_GetDeviceInfo(SAFEARRAY** pDevType, SAFEARRAY** pLedCount);
int MLAPI_GetLedInfo(BSTR type, DWORD index, BSTR* pName, SAFEARRAY** pLedStyles);
int MLAPI_GetLedColor(BSTR type, DWORD index, DWORD* R, DWORD* G, DWORD* B);
int MLAPI_GetLedStyle(BSTR type, DWORD index, BSTR* style);
int MLAPI_GetLedMaxBright(BSTR type, DWORD index, DWORD* maxLevel);
int MLAPI_GetLedBright(BSTR type, DWORD index, DWORD* currentLevel);
int MLAPI_GetLedMaxSpeed(BSTR type, DWORD index, DWORD* maxLevel);
int MLAPI_GetLedSpeed(BSTR type, DWORD index, DWORD* currentLevel);
int MLAPI_SetLedColor(BSTR type, DWORD index, DWORD R, DWORD G, DWORD B);
int MLAPI_SetLedStyle(BSTR type, DWORD index, BSTR style);
int MLAPI_SetLedBright(BSTR type, DWORD index, DWORD level);
int MLAPI_SetLedSpeed(BSTR type, DWORD index, DWORD level);
