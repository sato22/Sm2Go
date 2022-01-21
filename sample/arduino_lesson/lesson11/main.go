package main

import (
	"machine"
	"time"

	"Sm2Go/sample/arduino_lesson/lesson11/keypad4x4"
	"Sm2Go/sample/arduino_lesson/lesson11/pkg"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

func main() {
	led := machine.LED
	keypad := keypad4x4.NewDevice(machine.D9, machine.D8, machine.D7, machine.D6, machine.D5, machine.D4, machine.D3, machine.D2)

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	keypad.Configure()

	pkg.ConfigureDevice(led)
	pkg.ConfigureKeypad(keypad)

	pkg.ConfigureLog(logTest)

	for {
		pkg.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
