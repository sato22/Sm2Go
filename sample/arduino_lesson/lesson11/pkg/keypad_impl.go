// Please edit this file

package pkg

import (
	"time"
)

// interface
type Device interface {
	High()
	Low()
}

type Keypad interface {
	Configure()
	GetKey() string
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

// generated
func onEntry() {
	logger.Println("Success") // loggerを使った出力にする
	dev.High()
	entered = ""
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEmptyEntry() {
	dev.Low()
}

func offEmptyDo() {
	key := keypad4.GetKey()
	if key != "NoKeyPressed" {
		logger.Println("entered:" + key)
		entered += key
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
		logger.Println("Fail")
		logger.Println("Entered Passcode:" + entered)
		entered = ""
		return true
	} else {
		return false
	}
}
