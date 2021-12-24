// Please edit this file

package ledpkg

type Device interface {
	High()
	Low()
}

type Switch interface {
	Get() bool
}

var dev Device
var leftSwitch Switch

func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureLeftButton(s Switch) {
	leftSwitch = s
}

// generated
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

func time500msCond() bool {
	// time.Sleep(time.Millisecond * 1000)
	// return true
	return leftSwitch.Get()
}
