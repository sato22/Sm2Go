// Please run "gofmt -w /path/to/output.go"
// go run *.go

package main // write package name

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
			Wait_Entry()
			eod = Do
		}
		if eod == Do {
			Wait_Do()
			if push_button_Cond() {
				state = Run
				fmt.Println("State is changed: Wait to Run")
				eod = Entry
			}
		}
	case Run:
		if eod == Entry {
			Run_Entry()
			eod = Do
		}
		if eod == Do {
			Run_Do()
			if push_button_Cond() {
				state = Stop
				fmt.Println("State is changed: Run to Stop")
				eod = Entry
			}
		}
	}
}

func touch_display(state State, eod Eod) {
	switch state {
	case Stop:
		if eod == Entry {
			Stop_Entry()
			eod = Do
		}
		if eod == Do {
			Stop_Do()
			if touch_display_Cond() {
				state = Run
				fmt.Println("State is changed: Stop to Run")
				eod = Entry
			}
		}
	}
}

func main() {
	var Wait State = Wait
	// var Run State = Run
	// var Stop State = Stop
	var eod Eod = Entry
	var current_state State = Wait

	push_button(current_state, eod)
}
