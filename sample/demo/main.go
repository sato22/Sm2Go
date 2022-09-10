package main

import (
	"Sm2Go/sample/demo/ledbutton"
	"machine"
	"time"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

func main() {
	led := machine.D5
	leftButton := machine.D9
	rightButton := machine.D8

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	rightButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	ledbutton.ConfigureDevice(led)
	ledbutton.ConfigureLeftButton(leftButton)
	ledbutton.ConfigureRightButton(rightButton)

	ledbutton.ConfigureLog(logTest)

	for {
		ledbutton.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
