// Please edit this file

package ledbutton

type Device interface {
	High()
	Low()
}

type Switch interface {
	Get() bool
}

var dev Device
var leftSwitch Switch
var rightSwitch Switch

// var dev machine.Pin
// var leftSwitch machine.Pin
// var rightSwitch machine.Pin

func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureLeftButton(s Switch) {
	leftSwitch = s
}

func ConfigureRightButton(s Switch) {
	rightSwitch = s
}

// func ConfigureDevice(d machine.Pin) {
// 	dev = d
// }

// func ConfigureLeftButton(s machine.Pin) {
// 	leftSwitch = s
// }

// func ConfigureRightButton(s machine.Pin) {
// 	rightSwitch = s
// }

// generated
func onEntry() {
	dev.High()
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	dev.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func pushRightButtonCond() bool {
	return rightSwitch.Get()
}

func pushLeftButtonCond() bool {
	return leftSwitch.Get()
}
