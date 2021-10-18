// This is a test file for testing state transitions

package test

import (
//"fmt"
"testing"
"time"
)

func TestStateTrans(t *testing.T) {
	// ステートの遷移を確認可能
	go func() {
	for {
	time.Sleep(1 * time.Millisecond)
	task()
	}
	}()

	// goroutine
	for{}
}
