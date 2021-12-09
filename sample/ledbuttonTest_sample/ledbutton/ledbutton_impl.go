// Please edit this file

package main

// "machine"
// "log"

var led = &Led{"led", "High"}
var leftButton = &Button{"leftButton", false}
var rightButton = &Button{"rightButton", false}

func onEntry() {
	led.High()
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	led.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func pushRightButtonCond() bool {
	return rightButton.Get()
}

func pushLeftButtonCond() bool {
	return leftButton.Get()
}
