package main

import (
	"machine"
	"time"

	// "github.com/sato22/Sm2Go/sample/ledbuttonTest_sample/ledbutton"
	"Sm2Go/sample/ledbuttonTest_sample/ledbutton"
)

/*
どの部分でエラーが発生しているのか、問題の切り分けを行う
*/

func main() {
	led := machine.D5
	leftButton := machine.D9
	rightButton := machine.D8

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	rightButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		ledbutton.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
