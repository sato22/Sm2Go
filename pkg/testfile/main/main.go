
// main file

package main

import (
	"time"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

func main() {
	for {
modelTestTopStep()

		time.Sleep(time.Millisecond * 10)
	}
}
