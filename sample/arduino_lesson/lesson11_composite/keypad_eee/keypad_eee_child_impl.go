// Please edit this file

package keypad_eee

import (
	"strconv"
)

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
	key := Keypad4.GetKey()
	if key != NoKeyPressed {
		println("entered:", key)
		entered += strconv.Itoa(int(key))
	}
}

func offemptyExit() {
	// nothing to do
}

func keyFailedCond() bool {
	if passcode != entered {
		println("Fail")
		println("Entered Passcode: ", entered)
		entered = ""
		return true
	} else {
		return false
	}
}

func keyEnteredCond() bool {
	return len(passcode) == len(entered)
}
