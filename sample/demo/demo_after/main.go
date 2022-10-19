// main file

package main

import (
	"Sm2Go/sample/demo_tyukan/ledbutton"
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
	button := machine.D9

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	ledbutton.ConfigureDevice(led)
	ledbutton.ConfigureButton(button)

	for {
		ledbutton.Task()
		time.Sleep(time.Millisecond * 10)
	}
}
