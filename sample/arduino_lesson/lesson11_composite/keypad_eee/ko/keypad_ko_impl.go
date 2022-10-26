// Please edit this file

package ko

import (
	// "Sm2Go/sample/arduino_lesson/lesson11_composite/keypad_eee"
	"Sm2Go/sample/arduino_lesson/lesson11_composite/keypad_eee"
	"strconv"
)

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}



func offenteredEntry() {
	// nothing to do
}

func offenteredDo() {
	// nothing to do
}

func offenteredExit() {
	// nothing to do
}

func offemptyEntry() {
	// nothing to do
}

func offemptyDo() {
	key := keypad_eee.Keypad4.GetKey()
	if key != NoKeyPressed {
		println("entered:", key)
		keypad_eee.Entered += strconv.Itoa(int(key))
	}
}

func offemptyExit() {
	// nothing to do
}

func keyFailedCond() bool {
	if passcode != keypad_eee.Entered {
		println("Fail")
		println("Entered Passcode: ", keypad_eee.Entered)
		keypad_eee.Entered = ""
		return true
	} else {
		return false
	}
}

func keyEnteredCond() bool {
	return len(keypad_eee.Passcode) == len(keypad_eee.Entered)
}
