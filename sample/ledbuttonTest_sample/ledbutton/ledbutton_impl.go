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

func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureLeftButton(s Switch) {
	leftSwitch = s
}

func ConfigureRightButton(s Switch) {
	rightSwitch = s
}

// func deviceHigh(d Device) {
// 	d.High()
// }

// func deviceLow(d Device) {
// 	d.Low()
// }

// func switchGet(b Switch) bool {
// 	return b.Get()
// }

// generated

func onEntry() {
	// deviceHigh(dev)
	dev.High()
}

func onDo() {
	//nothing to do
}

func onExit() {
	// nothing to do
}

func offEntry() {
	// deviceLow(dev)
	dev.Low()
}

func offDo() {
	// nothing to do
}

func offExit() {
	// nothing to do
}

func pushRightButtonCond() bool {
	// return switchGet(rightSwitch)
	return rightSwitch.Get()
}

func pushLeftButtonCond() bool {
	// return switchGet(leftSwitch)
	return leftSwitch.Get()
}
