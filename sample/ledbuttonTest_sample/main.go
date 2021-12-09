package main

import (
	"machine"
	"time"

	"github.com/sato22/Sm2Go/sample/ledbuttonTest_sample/ledbutton"
)

func main() {
	led := machine.D5
	leftButton := machine.D9
	rightButton := machine.D8

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	rightButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	ledbutton.ConfigureDevice(&led)
	ledbutton.ConfigureLeftButton(&leftButton)
	ledbutton.ConfigureRightButton(&rightButton)

	for {
		ledbutton.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
