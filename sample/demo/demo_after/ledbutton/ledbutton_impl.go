// Please edit this file

package ledbutton

// package name to import

// interface
type Device interface {
	High()
	Low()
}

type Switch interface {
	Get() bool
}

var dev Device
var button Switch

// configure
func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureButton(s Switch) {
	button = s
}

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

var oldValue = true

func onEntry() {
	dev.High()
}

func onDo() {
	// nothing to do
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

func pushButtonCond() bool {
	// Please write the conditions under which a state transitions

	if button.Get() == false && oldValue == true {
		// 押した瞬間
		oldValue = button.Get()
		return true
	}

	oldValue = button.Get()

	return false
}
