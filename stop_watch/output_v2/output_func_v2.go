// Please run "gofmt -w /path/to/output.go"
// go run *.go

package main // write package name

import (
	// write about import
	"fmt"
)

func Wait_Entry() {
	// Processing performed in the "Entry" state of the "Wait" state.
	fmt.Println("Wait_Entry() is executed.")
}

func Run_Entry() {
	// Processing performed in the "Entry" state of the "Run" state.
}

func Stop_Entry() {
	// Processing performed in the "Entry" state of the "Stop" state.
}

func Wait_Do() {
	// Processing performed in the "Do" state of the "Wait" state.
}

func Run_Do() {
	// Processing performed in the "Do" state of the "Run" state.
}

func Stop_Do() {
	// Processing performed in the "Do" state of the "Stop" state.
}

func push_button_Cond() bool {
	// Please write the conditions under which a state transitions
	return true
}

func touch_display_Cond() bool {
	// Please write the conditions under which a state transitions
	return true
}
