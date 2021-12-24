package main

import (
	"machine"
	"time"

	"Sm2Go/sample/ledbutton_sample/ledbutton/ledpkg"
)

func main() {
	led := machine.D5
	leftButton := machine.D9
	rightButton := machine.D8

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	rightButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	ledpkg.ConfigureDevice(led)
	ledpkg.ConfigureLeftButton(leftButton)
	ledpkg.ConfigureRightButton(rightButton)

	// count := 0

	for {
		ledpkg.Task()
		time.Sleep(time.Millisecond * 10)
		// count++
		// println("main.go print test", count)
		// fmt.Println("main.go print test")
	}
}
