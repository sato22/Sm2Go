// Please edit this file

package main

import (
	// "machine"
	// "log"
)

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
