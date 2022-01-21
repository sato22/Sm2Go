package main

import (
	"machine"
	"time"

	"Sm2Go/sample/led_sample/led/ledpkg"
)

func main() {
	led := machine.D5
	leftButton := machine.D9

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	ledpkg.ConfigureDevice(led)
	ledpkg.ConfigureLeftButton(leftButton)

	for {
		time.Sleep(1 * time.Millisecond)
		ledpkg.Task()
	}
}
