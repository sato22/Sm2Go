package main

import (
	"machine"
	"sync"
	"time"
)

func main() {
	led = machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			task()
		}
		wg.Done()
	}()
	wg.Wait()
}
