// Please edit this file
// go run *.go

package main

import (
	"fmt"
	"time"
)

func Wait_Entry() {
	// nothing to do
}

func Wait_Do() {
	// nothing to do
}

func Wait_Exit() {
	// nothing to do
}

func Run_Entry() {
	// nothing to do
}

func Run_Do() {
	// nothing to do
}

func Run_Exit() {
	// nothing to do
}

func Stop_Entry() {
	// nothing to do
}

func Stop_Do() {
	// nothing to do
}

func Stop_Exit() {
	// nothing to do
}

func push_button_Cond() bool {
	// Please write the conditions under which a state transitions
	return false
}

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			task1()
		}
	}()
	for {
		fmt.Scan(&input)
		if input == "q" {
			fmt.Println("quit")
			break
		}
	}
}
