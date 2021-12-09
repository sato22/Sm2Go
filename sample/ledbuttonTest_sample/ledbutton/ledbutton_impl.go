// Please edit this file

package ledbutton

type Device interface {
	High()
	Low()
}

type Switch interface {
	Get() bool
}

func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureLeftButton(s Switch) {
	leftSwitch = s
}

func ConfigureRightButton(s Switch) {
	rightSwitch = s
}

var dev Device
var leftSwitch Switch
var rightSwitch Switch

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
