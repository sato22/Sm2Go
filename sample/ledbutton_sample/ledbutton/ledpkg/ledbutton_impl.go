// Please edit this file

package ledpkg

// interface
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

// configure
func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureLeftButton(s Switch) {
	leftSwitch = s
}

func ConfigureRightButton(s Switch) {
	rightSwitch = s
}

// log.Println()
type DebugPrint interface {
	PrintLog(string)
}

var debug DebugPrint

func ConfigureLog(p DebugPrint) {
	debug = p
}

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
	return !rightSwitch.Get()
}

func pushLeftButtonCond() bool {
	return !leftSwitch.Get()
}
