// Please run "gofmt -w /path/to/output.go"
// go run *.go

package main // write package name

import (
	// write about import
	"fmt"
)

func Heat_Entry() {
	// Processing performed in the "Entry" state of the "Heat" state.
	fmt.Println("Heat_Entry() is executed")
}

func Unheat_Entry() {
	// Processing performed in the "Entry" state of the "Unheat" state.
}

func Stop_Entry() {
	// Processing performed in the "Entry" state of the "Stop" state.
}

func Heat_Do() {
	// Processing performed in the "Do" state of the "Heat" state.
	fmt.Println("Heat_Do() is executed")
}

func Unheat_Do() {
	// Processing performed in the "Do" state of the "Unheat" state.
}

func Stop_Do() {
	// Processing performed in the "Do" state of the "Stop" state.
}

func higher_limit_Cond() bool {
	// Please write the conditions under which a state transitions
	return true
}

func lower_limit_Cond() bool {
	// Please write the conditions under which a state transitions
	return true
}

func power_off_Cond() bool {
	// Please write the conditions under which a state transitions
	fmt.Println("power_off_Cond() is executed")
	return true
}
