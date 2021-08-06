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

var eod Eod
var current_state State

func state_transition() {
	switch current_state {
	case Wait:
		if eod == Entry {
			Wait_Entry()
			eod = Do
		}
		if eod == Do {
			Wait_Do()
			if push_button_Cond() {
				current_state = Run
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
				current_state = Stop
				fmt.Println("State is changed: Run to Stop")
				eod = Entry
			}
		}

	case Stop:
		if eod == Entry {
			Stop_Entry()
			eod = Do
		}
		if eod == Do {
			Stop_Do()
			if push_button_Cond() {
				current_state = Wait
				fmt.Println("State is changed: Stop to Wait")
				eod = Entry
			}
		}
	}
}

func init() {
	current_state = Wait
	eod = Entry
}
