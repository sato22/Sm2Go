package modelTest


// This is a test file for testing state transitions

import (
	"log"
	"testing"
	"time"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	log.Println(debstr)
}
func TopTestDevice(t *testing.T) {
	env := sm2go.NewTestEnv() // TestEnv構造体

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
TopStep()

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
