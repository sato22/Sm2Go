// Please edit this file

package keypad_eee

import (
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

var Dev Device
var Keypad4 Keypad

// configure
func ConfigureDevice(d Device) {
	Dev = d
}

func ConfigureKeypad(k Keypad) {
	Keypad4 = k
}

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

// 別のパッケージで定義
var Passcode string = "0000"
var Entered string = ""

const NoKeyPressed = 255

func onEntry() {
	logger.Println("Success") // loggerを使った出力にする
	Dev.High()
	Entered = ""
}

func onDo() {
	// nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	Dev.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func time3secCond() bool {
	time.Sleep(time.Second * 3)
	return true
}

func correctKeyCond() bool {
	return Passcode == Entered
}
