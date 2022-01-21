// Please edit this file

package pkg

import (
	"strconv"
	"time"
)

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

// debuglog
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

// generated
func onEntry() {
	println("Success")
	dev.High()
	entered = ""
}

func onDo() {
	//nothing to do
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
	currentSmall = OffEmpty
}

func offEmptyEntry() {
	// nothing to do
}

func offEmptyDo() {
	key := keypad4.GetKey()
	if key != NoKeyPressed {
		println("entered:", key)
		entered += strconv.Itoa(int(key))
	}
}

func offEmptyExit() {
	// nothing to do
}

func offEnteredEntry() {
	// nothing to do
}

func offEnteredDo() {
	// nothing to do
}

func offEnteredExit() {
	// nothing to do
}

func correctKeyCond() bool {
	return passcode == entered
}

func time3secCond() bool {
	time.Sleep(time.Second * 3)
	return true
}

func keyEnteredCond() bool {
	return len(passcode) == len(entered)
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
