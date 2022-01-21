// This is a test file for testing state transitions

package keypad

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
// goroutine (base.go Task())
func TestStateTrans(t *testing.T) {
var wg sync.WaitGroup
wg.Add(1)
go func() {
for {
time.Sleep(1 * time.Millisecond)
Task()
}
}()
wg.Wait()
}

func TestDevice(t *testing.T) {
var wg sync.WaitGroup
wg.Add(1)
// goroutine(base.go Task())
go func() {
for {
time.Sleep(50 * time.Millisecond)
Task()
}
// goroutine(user operation)
wg.Done()
}()
wg.Wait()
}

