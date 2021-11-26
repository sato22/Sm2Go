// Please edit this file

package ledbutton

import (
	// "machine"
	"log"
)

func onEntry() {
	// machine.D5.High()
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	// machine.D5.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func pushRightButtonCond() bool {
	// return !machine.D8.Get()
	log.Println("rightButton.Get() =", rightButton.Get())
	return rightButton.Get()
}

func pushLeftButtonCond() bool {
	// return !machine.D9.Get()
	return leftButton.Get()
}
