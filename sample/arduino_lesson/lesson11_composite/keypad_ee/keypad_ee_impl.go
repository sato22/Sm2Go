// Please edit this file

package keypad_easy

import (
	"strconv"
	"time"
)

// package name to import

// interface
type Device interface {
	High()
	Low()
}

type Keypad interface {
	Configure()
	GetKey() uint8
}

var dev Device
var keypad4 Keypad

// configure
func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureKeypad(k Keypad) {
	keypad4 = k
}

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

var passcode string = "0000"
var entered string = ""

const NoKeyPressed = 255

func onEntry() {
	logger.Println("Success") // loggerを使った出力にする
	dev.High()
	entered = ""
}

func onDo() {
	// nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	dev.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// これはどういうこと？
	currentSmall = OffEmpty
}

func time3secCond() bool {
	time.Sleep(time.Second * 3)
	return true
}

func correctKeyCond() bool {
	return passcode == entered
}

// // Please edit this file

// package keypad_easy

// import (
// // package name to import
// )

// type DebugLogger interface {
// Println(string)
// }

// var logger DebugLogger

// func ConfigureLog(p DebugLogger) {
// logger = p
// }

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
	key := keypad4.GetKey()
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
	// fmt.Printf("%t\n", len(passcode) == len(entered))
	// fmt.Println("len(entered)", len(entered))
	return len(passcode) == len(entered)
}
