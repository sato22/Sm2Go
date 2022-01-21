// main file

package main

import (
"machine"
"time"
)

type DebugStruct struct{}
var logTest = DebugStruct{}
func (l DebugStruct) Println(debstr string) {
println(debstr)
}
func main() {
for {
Task()
time.sleep(time.Millisecond * 10)
}
}

