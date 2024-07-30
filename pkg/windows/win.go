package windows

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

func MessageBox(Hwnd uintptr, Caption, Title string, Flags uint) int {
	Ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(Hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(Caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(Title))),
		uintptr(Flags))
	return int(Ret)
}

func MsgBox(Title, Caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, Caption, Title, MB_OK)
}

func SetWindowSize(x string, y string) {
	Cmd := exec.Command("cmd.exe", "/c", fmt.Sprintf("mode con: cols=%s lines=%s", x, y))
	Cmd.Stdout = os.Stdout
	Cmd.Run()
}

func SetConsoleTitle(Title string) error {
	Kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}
	defer syscall.FreeLibrary(Kernel32)
	SetConsoleTitle, err := syscall.GetProcAddress(Kernel32, "SetConsoleTitleW")
	if err != nil {
		return err
	}
	TitlePtr, err := syscall.UTF16PtrFromString(Title)
	if err != nil {
		return err
	}
	_, _, CallErr := syscall.Syscall(SetConsoleTitle, 1, uintptr(unsafe.Pointer(TitlePtr)), 0, 0)
	if CallErr != 0 {
		return CallErr
	}
	return nil
}
