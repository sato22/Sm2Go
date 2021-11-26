package main

import (
	"machine"
	"time"
	"github.com/sato22/Sm2Go/sample/ledbuttonTest_sample/ledbutton"
)

const (
	led  = machine.D5
	left = machine.D9
	right = machine.D8
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	left.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	right.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for {
		task()
		time.Sleep(time.Millisecond * 10)
	}
}
