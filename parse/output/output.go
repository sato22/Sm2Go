// Please run "gofmt -w /path/to/output.go"

package parser // input package name

import (
	// write about import
	"fmt"
)

type State int

const (
	Wait State = iota
	Run
	Stop
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

func push_button(state State, eod Eod) {
	switch state {
	case Wait:
		if eod == Entry {
			wait_test_entryProg
			eod = Do
		}
		if eod == Do {
			wait_test_doProg
			if ecrobot_get_touch_sensor(NXT_PORT_TOUCH) == 1 {
				state = Run
				fmt.Println("State is changed: Wait to Run ")
				eod = Entry
			}
		}
	case Run:
		if eod == Entry {
			run_test_entryProg
			eod = Do
		}
		if eod == Do {
			run_test_doProg
			if ecrobot_get_touch_sensor(NXT_PORT_TOUCH) == 1 {
				state = Stop
				fmt.Println("State is changed: Run to Stop ")
				eod = Entry
			}
		}
	}
}

func touch_display(state State, eod Eod) {
	switch state {
	case Stop:
		if eod == Entry {
			stop_test_entryProg
			eod = Do
		}
		if eod == Do {
			stop_test_doProg
			if get_touch_sensor(NXT_PORT_TOUCH) == 1 {
				state = Run
				fmt.Println("State is changed: Stop to Run ")
				eod = Entry
			}
		}
	}
}
