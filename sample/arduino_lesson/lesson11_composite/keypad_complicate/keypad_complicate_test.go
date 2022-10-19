// This is a test file for testing state transitions

package keypad_complicate

import (
"log"
"testing"
"time"
"sync"
)

type DebugStruct struct{}
var logTest = DebugStruct{}
func (l DebugStruct) Println(debstr string) {
log.Println(debstr)
}
func TestDevice(t *testing.T) {
env := sm2go.NewTestEnv() // TestEnv構造体

// goroutine(base.go Task())
env.Add(sm2go.Continue, func() {
for {
time.Sleep(10 * time.Millisecond)
Task()
}
},
)
// goroutine(user operation)
env.Add(sm2go.Done, func() {
},
)

env.Set(1)
env.Go()
}

// This is a test file for testing state transitions

package keypad_complicate

import (
"log"
"testing"
"time"
"sync"
)

type DebugStruct struct{}
var logTest = DebugStruct{}
func (l DebugStruct) Println(debstr string) {
log.Println(debstr)
}
func TestDevice(t *testing.T) {
env := sm2go.NewTestEnv() // TestEnv構造体

// goroutine(base.go Task())
env.Add(sm2go.Continue, func() {
for {
time.Sleep(10 * time.Millisecond)
Task()
}
},
)
// goroutine(user operation)
env.Add(sm2go.Done, func() {
},
)

env.Set(1)
env.Go()
}

