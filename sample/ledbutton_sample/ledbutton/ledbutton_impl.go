// Please edit this file

package main

import (
	"machine"
)

func onEntry() {
	machine.D5.High()
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	machine.D5.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func pushRightButtonCond() bool {
	if machine.D8.Get(){
		return false
	}
	return true
}

func pushLeftButtonCond() bool {
	if machine.D9.Get(){
		return false
	}
	return true
}
