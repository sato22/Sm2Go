package modelTest


// This is a test file for testing state transitions

import (
	"log"
	"testing"
	"time"
)
func GrandchildTestDevice(t *testing.T) {
	env := sm2go.NewTestEnv() // TestEnv構造体

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
GrandchildStep()

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
