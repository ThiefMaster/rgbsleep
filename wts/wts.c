// taken from https://github.com/bbigras/go-windows-session-notifications

#define _WIN32_WINNT 0x0A00 /* win10 */

// missing in wtsapi32.h
typedef enum _WTS_VIRTUAL_CLASS {
    WTSVirtualClientData,
    WTSVirtualFileHandle
} WTS_VIRTUAL_CLASS;

#include <windows.h>
#include <wtsapi32.h>
#include "_cgo_export.h"

LRESULT CALLBACK WndProc(HWND hWnd, UINT message, WPARAM wParam, LPARAM lParam)
{
    int wmId, wmEvent;
    switch (message)
    {
    //case WM_QUERYENDSESSION:
    //    relayMessage(message, lParam);
    //    break;
    case WM_WTSSESSION_CHANGE:
        relayMessage(message, wParam);
        break;
    default:
        return DefWindowProc(hWnd, message, wParam, lParam);
    }
    return 0;
}

DWORD WINAPI WatchSessionNotifications(LPVOID lpParam)
{
    WNDCLASS wc;
    HWND hwnd;
    MSG msg;

    char const *lpClassName = "classWatchSessionNotifications";

    wc.lpfnWndProc = WndProc;
    wc.lpszClassName=lpClassName;

    if (!RegisterClass(&wc)) {
        return 0;
    }

    hwnd = CreateWindow(lpClassName,
                        lpClassName,
                        WS_OVERLAPPEDWINDOW,
                        CW_USEDEFAULT, CW_USEDEFAULT, 100, 100,
                        NULL, NULL, NULL, NULL);

    if (!hwnd) {
        return 0;
    }

    UpdateWindow(hwnd);
    WTSRegisterSessionNotification(hwnd, NOTIFY_FOR_THIS_SESSION);

    while (GetMessage(&msg,NULL,0,0) > 0) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }
}

void Stop(HANDLE hndl) {
    // TODO: Figure out how to ask the thread to quit by itself
    // PostThreadMessage(hndl, WM_QUIT, 0, 0);
    // WaitForSingleObject(hndl, 5000);
    TerminateThread(hndl, 0);
}

HANDLE Start() {
    DWORD lpThreadId, lpParameter = 1;
    HANDLE hThread;

    hThread = CreateThread(
                  NULL,
                  0,
                  WatchSessionNotifications,
                  &lpParameter,
                  0,
                  &lpThreadId);

    if (hThread == NULL) {
        return 0;
    } else {
        return hThread;
    }
}
