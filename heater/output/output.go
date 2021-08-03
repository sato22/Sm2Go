// Please run "gofmt -w /path/to/output.go"
// go run *.go

package main // write package name

import (
	// write about import
	"fmt"
)

type State int

const (
	Heat State = iota
	Unheat
	Stop
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

func higher_limit(state State, eod Eod) {
	switch state {
	case Heat:
		if eod == Entry {
			Heat_Entry()
			eod = Do
		}
		if eod == Do {
			Heat_Do()
			if higher_limit_Cond() {
				state = Unheat
				fmt.Println("State is changed: Heat to Unheat")
				eod = Entry
			}
		}
	}
}

func lower_limit(state State, eod Eod) {
	switch state {
	case Unheat:
		if eod == Entry {
			Unheat_Entry()
			eod = Do
		}
		if eod == Do {
			Unheat_Do()
			if lower_limit_Cond() {
				state = Heat
				fmt.Println("State is changed: Unheat to Heat")
				eod = Entry
			}
		}
	}
}

func power_off(state State, eod Eod) {
	switch state {
	case Heat:
		if eod == Entry {
			Heat_Entry()
			eod = Do
		}
		if eod == Do {
			Heat_Do()
			if power_off_Cond() {
				state = Stop
				fmt.Println("State is changed: Heat to Stop")
				eod = Entry
			}
		}
	case Unheat:
		if eod == Entry {
			Unheat_Entry()
			eod = Do
		}
		if eod == Do {
			Unheat_Do()
			if power_off_Cond() {
				state = Stop
				fmt.Println("State is changed: Unheat to Stop")
				eod = Entry
			}
		}
	}
}

func main() {
	var Heat State = Heat
	// var Unheat State = Unheat
	// var Stop State = Stop
	var eod Eod = Entry

	var current_state State = Heat

	power_off(current_state, eod)
}
