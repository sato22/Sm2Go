package main

import (
	"machine"
	"time"
	"github.com/sato22/Sm2Go/sample/ledbuttonTest_sample/ledbutton"
)

const (
	led  = machine.D5
	leftButton = machine.D9
	rightButton = machine.D8
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	rightButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for {
		ledbutton.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
