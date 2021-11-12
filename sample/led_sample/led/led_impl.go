// Please edit this file

package main

import (
	"machine"
	"time"
)

func onEntry() {
	// nothing to do
}

func onDo() {
	machine.LED.High()
}

func onExit() {
	// nothing to do
}

func offEntry() {
	// nothing to do
}

func offDo() {
	machine.LED.Low()
}

func offExit() {
	// nothing to do
}

func time500msCond() bool {
	time.Sleep(time.Millisecond * 1000)
	return true
}
