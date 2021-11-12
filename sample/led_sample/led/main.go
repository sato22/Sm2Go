package main

import (
	"machine"
	//"sync"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		time.Sleep(1 * time.Millisecond)
		task()
	}
}
