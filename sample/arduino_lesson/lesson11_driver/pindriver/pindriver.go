package pindriver

// import "machine"

const (
	PinLow uint8 = iota
	PinHigh
)

type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
	PinInputPullup
)

type PinConfig struct {
	Mode PinMode
}

type Pinif interface {
	High()
	Low()
	Set(bool)
	Get() bool
	Configure(config interface{})
	// ConfigureTest(config interface{})
}

// func (p machine.Pin) ConfigureTest(config interface{}) {
// 	s, _ := config.(machine.PinConfig)
// 	p.Configure(s)
// }

// func (p pindriver.Pinif) ConfigureTest(config interface{}) {
// 	s, _ := config.(machine.PinConfig)
// 	p.Configure(s)
// }
